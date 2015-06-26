package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/swag"
)

// GenerateDefinition generates a model file for a schema defintion.
func GenerateDefinition(modelNames []string, includeModel, includeValidator bool, opts GenOpts) error {
	// Load the spec
	specPath, specDoc, err := loadSpec(opts.Spec)
	if err != nil {
		return err
	}

	if len(modelNames) == 0 {
		for k := range specDoc.Spec().Definitions {
			modelNames = append(modelNames, k)
		}
	}

	for _, modelName := range modelNames {
		// lookup schema
		model, ok := specDoc.Spec().Definitions[modelName]
		if !ok {
			return fmt.Errorf("model %q not found in definitions in %s", modelName, specPath)
		}

		// generate files
		generator := definitionGenerator{
			Name:             modelName,
			Model:            model,
			SpecDoc:          specDoc,
			Target:           filepath.Join(opts.Target, opts.ModelPackage),
			IncludeModel:     includeModel,
			IncludeValidator: includeValidator,
			DumpData:         opts.DumpData,
		}

		if err := generator.Generate(); err != nil {
			return err
		}
	}

	return nil
}

type definitionGenerator struct {
	Name             string
	Model            spec.Schema
	SpecDoc          *spec.Document
	Target           string
	IncludeModel     bool
	IncludeValidator bool
	Data             interface{}
	DumpData         bool
}

func (m *definitionGenerator) Generate() error {
	mod, err := makeGenDefinition(m.Name, m.Target, m.Model, m.SpecDoc)
	if err != nil {
		return err
	}
	if m.DumpData {
		bb, _ := json.MarshalIndent(swag.ToDynamicJSON(mod), "", " ")
		fmt.Fprintln(os.Stdout, string(bb))
		return nil
	}

	mod.IncludeValidator = m.IncludeValidator
	m.Data = mod

	if m.IncludeModel {
		if err := m.generateModel(); err != nil {
			return fmt.Errorf("model: %s", err)
		}
	}
	log.Println("generated model", m.Name)

	return nil
}

func (m *definitionGenerator) generateModel() error {
	buf := bytes.NewBuffer(nil)

	if err := modelTemplate.Execute(buf, m.Data); err != nil {
		return err
	}
	log.Println("rendered model template:", m.Name)

	return writeToFile(m.Target, m.Name, buf.Bytes())
}

func makeGenDefinition(name, pkg string, schema spec.Schema, specDoc *spec.Document) (*GenDefinition, error) {
	receiver := "m"
	resolver := &typeResolver{
		ModelsPackage: "",
		ModelName:     name,
		Doc:           specDoc,
	}
	pg := schemaGenContext{
		Path:         "",
		Name:         name,
		Receiver:     receiver,
		IndexVar:     "i",
		ValueExpr:    receiver,
		Schema:       schema,
		Required:     false,
		TypeResolver: resolver,
		Named:        true,
		ExtraSchemas: make(map[string]GenSchema),
	}
	if err := pg.makeGenSchema(); err != nil {
		return nil, err
	}

	var defaultImports []string
	if pg.GenSchema.HasValidations {
		defaultImports = []string{
			"github.com/go-swagger/go-swagger/errors",
			"github.com/go-swagger/go-swagger/strfmt",
			"github.com/go-swagger/go-swagger/httpkit/validate",
		}
	}
	var extras []GenSchema
	for _, v := range pg.ExtraSchemas {
		extras = append(extras, v)
	}

	return &GenDefinition{
		Package:        filepath.Base(pkg),
		GenSchema:      pg.GenSchema,
		DependsOn:      pg.Dependencies,
		DefaultImports: defaultImports,
		ExtraSchemas:   extras,
	}, nil
}

// GenDefinition contains all the properties to generate a
// defintion from a swagger spec
type GenDefinition struct {
	GenSchema
	Package          string
	Imports          map[string]string
	DefaultImports   []string
	ExtraSchemas     []GenSchema
	DependsOn        []string
	IncludeValidator bool
}

// GenSchemaList is a list of schemas for generation.
//
// It can be sorted by name to get a stable struct layout for
// version control and such
type GenSchemaList []GenSchema

func (g GenSchemaList) Len() int           { return len(g) }
func (g GenSchemaList) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g GenSchemaList) Less(i, j int) bool { return g[i].Name < g[j].Name }

type schemaGenContext struct {
	Path               string
	Name               string
	ParamName          string
	Accessor           string
	Receiver           string
	IndexVar           string
	KeyVar             string
	ValueExpr          string
	Schema             spec.Schema
	Required           bool
	AdditionalProperty bool
	TypeResolver       *typeResolver
	Untyped            bool
	Named              bool
	Index              int

	GenSchema    GenSchema
	Dependencies []string
	ExtraSchemas map[string]GenSchema
}

func (sg *schemaGenContext) NewSliceBranch(schema *spec.Schema) *schemaGenContext {
	pg := sg.shallowClone()
	indexVar := pg.IndexVar
	if pg.Path == "" {
		pg.Path = "strconv.Itoa(" + indexVar + ")"
	} else {
		pg.Path = pg.Path + "+ \".\" + strconv.Itoa(" + indexVar + ")"
	}
	pg.IndexVar = indexVar + "i"
	pg.ValueExpr = pg.ValueExpr + "[" + indexVar + "]"
	pg.Schema = *schema
	pg.Required = false

	// when this is an anonymous complex object, this needs to become a ref
	return pg
}

func (sg *schemaGenContext) NewAdditionalItems(schema *spec.Schema) *schemaGenContext {
	pg := sg.shallowClone()
	indexVar := pg.IndexVar
	pg.Name = sg.Name + " items"
	itemsLen := 0
	if sg.Schema.Items != nil {
		itemsLen = sg.Schema.Items.Len()
	}
	var mod string
	if itemsLen > 0 {
		mod = "+" + strconv.Itoa(itemsLen)
	}
	if pg.Path == "" {
		pg.Path = "strconv.Itoa(" + indexVar + mod + ")"
	} else {
		pg.Path = pg.Path + "+ \".\" + strconv.Itoa(" + indexVar + mod + ")"
	}
	pg.IndexVar = indexVar
	pg.ValueExpr = sg.ValueExpr + "." + swag.ToGoName(sg.Name) + "Items[" + indexVar + "]"
	pg.Schema = spec.Schema{}
	if schema != nil {
		pg.Schema = *schema
	}
	pg.Required = false
	return pg
}

func (sg *schemaGenContext) NewTupleElement(schema *spec.Schema, index int) *schemaGenContext {
	pg := sg.shallowClone()
	if pg.Path == "" {
		pg.Path = "\"" + strconv.Itoa(index) + "\""
	} else {
		pg.Path = pg.Path + "+ \".\"+\"" + strconv.Itoa(index) + "\""
	}
	pg.ValueExpr = pg.ValueExpr + ".P" + strconv.Itoa(index)
	pg.Required = true
	pg.Schema = *schema
	return pg
}

func (sg *schemaGenContext) NewStructBranch(name string, schema spec.Schema) *schemaGenContext {
	pg := sg.shallowClone()
	if sg.Path == "" {
		pg.Path = fmt.Sprintf("%q", name)
	} else {
		pg.Path = pg.Path + "+\".\"+" + fmt.Sprintf("%q", name)
	}
	pg.Name = name
	pg.ValueExpr = pg.ValueExpr + "." + swag.ToGoName(name)
	pg.Schema = schema
	for _, fn := range sg.Schema.Required {
		if name == fn {
			pg.Required = true
			break
		}
	}
	return pg
}

func (sg *schemaGenContext) shallowClone() *schemaGenContext {
	pg := new(schemaGenContext)
	*pg = *sg
	pg.GenSchema = GenSchema{}
	pg.Dependencies = nil
	//pg.ExtraSchemas = make(map[string]GenSchema)
	pg.Named = false
	pg.Index = 0
	return pg
}

func (sg *schemaGenContext) NewCompositionBranch(schema spec.Schema, index int) *schemaGenContext {
	pg := sg.shallowClone()
	pg.Schema = schema
	pg.Name = "AO" + strconv.Itoa(index)
	if sg.Name != sg.TypeResolver.ModelName {
		pg.Name = sg.Name + pg.Name
	}
	pg.Index = index
	return pg
}

func (sg *schemaGenContext) NewAdditionalProperty(schema spec.Schema) *schemaGenContext {
	pg := sg.shallowClone()
	pg.Schema = schema
	if pg.KeyVar == "" {
		pg.ValueExpr = ""
	}
	pg.KeyVar += "k"
	pg.ValueExpr += "v"
	pg.Path = pg.KeyVar
	pg.GenSchema.Suffix = "Value"
	if sg.Path != "" {
		pg.Path = sg.Path + "+\".\"+" + pg.KeyVar
	}
	return pg
}

func (sg *schemaGenContext) schemaValidations() sharedValidations {
	model := sg.Schema

	isRequired := sg.Required
	if sg.Schema.Default != nil || sg.Schema.ReadOnly {
		isRequired = false
	}
	hasNumberValidation := model.Maximum != nil || model.Minimum != nil || model.MultipleOf != nil
	hasStringValidation := model.MaxLength != nil || model.MinLength != nil || model.Pattern != ""
	hasSliceValidations := model.MaxItems != nil || model.MinItems != nil || model.UniqueItems
	hasValidations := isRequired || hasNumberValidation || hasStringValidation || hasSliceValidations

	//var enum string
	if len(sg.Schema.Enum) > 0 {
		hasValidations = true
	}

	return sharedValidations{
		Required:            sg.Required,
		Maximum:             model.Maximum,
		ExclusiveMaximum:    model.ExclusiveMaximum,
		Minimum:             model.Minimum,
		ExclusiveMinimum:    model.ExclusiveMinimum,
		MaxLength:           model.MaxLength,
		MinLength:           model.MinLength,
		Pattern:             model.Pattern,
		MaxItems:            model.MaxItems,
		MinItems:            model.MinItems,
		UniqueItems:         model.UniqueItems,
		MultipleOf:          model.MultipleOf,
		Enum:                sg.Schema.Enum,
		HasValidations:      hasValidations,
		HasSliceValidations: hasSliceValidations,
	}
}
func (sg *schemaGenContext) MergeResult(other *schemaGenContext) {
	if other.GenSchema.AdditionalProperties != nil && other.GenSchema.AdditionalProperties.HasValidations {
		sg.GenSchema.HasValidations = true
	}
	if other.GenSchema.HasValidations {
		sg.GenSchema.HasValidations = other.GenSchema.HasValidations
	}
	sg.Dependencies = append(sg.Dependencies, other.Dependencies...)
	for k, v := range other.ExtraSchemas {
		sg.ExtraSchemas[k] = v
	}
}

func (sg *schemaGenContext) buildProperties() error {
	for k, v := range sg.Schema.Properties {
		emprop := sg.NewStructBranch(k, v)
		if err := emprop.makeGenSchema(); err != nil {
			return err
		}
		sg.MergeResult(emprop)
		sg.GenSchema.Properties = append(sg.GenSchema.Properties, emprop.GenSchema)
	}
	sort.Sort(sg.GenSchema.Properties)
	return nil
}

func (sg *schemaGenContext) buildAllOf() error {
	for i, sch := range sg.Schema.AllOf {
		comprop := sg.NewCompositionBranch(sch, i)
		if err := comprop.makeGenSchema(); err != nil {
			return err
		}
		sg.MergeResult(comprop)
		sg.GenSchema.AllOf = append(sg.GenSchema.AllOf, comprop.GenSchema)
	}
	return nil
}

type mapType struct {
	Type     *spec.Schema
	Next     *mapType
	Previous *mapType
	Context  *schemaGenContext
	NewObj   *schemaGenContext
}

func (sg *schemaGenContext) buildAdditionalProperties() error {
	if sg.Schema.AdditionalProperties != nil {
		addp := *sg.Schema.AdditionalProperties
		wantsAdditional := addp.Allows || addp.Schema != nil
		sg.GenSchema.HasAdditionalProperties = wantsAdditional
		// log.Printf("%s (complex: %t, map: %t, hasAdditional: %t)", sg.Name, sg.GenSchema.IsComplexObject, sg.GenSchema.IsMap, wantsAdditional)
		// b, _ := json.MarshalIndent(sg.Schema, "", "  ")
		// fmt.Println(string(b))
		// flag swap
		if sg.GenSchema.IsComplexObject {
			sg.GenSchema.IsAdditionalProperties = true
			sg.GenSchema.IsComplexObject = false
			sg.GenSchema.IsMap = false
		}

		if addp.Schema != nil {
			if !sg.GenSchema.IsMap && (sg.GenSchema.IsAdditionalProperties && sg.Named) {
				// tpe := sg.GenSchema
				// log.Printf("additional properties for definition (complex: %t, anonymous: %t)", tpe.IsComplexObject, tpe.IsAnonymous)
				sg.GenSchema.ValueExpression += "." + sg.GenSchema.Name
				comprop := sg.NewAdditionalProperty(*addp.Schema)
				if err := comprop.makeGenSchema(); err != nil {
					return err
				}
				sg.MergeResult(comprop)
				sg.GenSchema.AdditionalProperties = &comprop.GenSchema
				return nil
			}

			if sg.GenSchema.IsMap && wantsAdditional {
				// find out how deep this rabbit hole goes
				// descend, unwind and rewrite
				stack := &mapType{
					Type:    sg.Schema.AdditionalProperties.Schema,
					Context: sg.NewAdditionalProperty(*sg.Schema.AdditionalProperties.Schema),
				}
				lastMapSchema := stack
				for lastMapSchema.Type.AdditionalProperties != nil && (lastMapSchema.Type.AdditionalProperties.Allows || lastMapSchema.Type.AdditionalProperties.Schema != nil) {
					tpe, err := sg.TypeResolver.ResolveSchema(lastMapSchema.Type.AdditionalProperties.Schema, true)
					if err != nil {
						return err
					}
					// log.Printf("%s context (complex: %t, anonymous: %t, map: %t)", sg.Name, sg.GenSchema.IsComplexObject, sg.GenSchema.IsAnonymous, sg.GenSchema.IsMap)
					// log.Printf("%s additional properties for level down (complex: %t, anonymous: %t, map: %t)", tpe.GoType, tpe.IsComplexObject, tpe.IsAnonymous, tpe.IsMap)
					if !tpe.IsMap {
						if tpe.IsComplexObject && tpe.IsAnonymous {
							nw := sg.makeNewStruct(sg.Name+" Anon", *lastMapSchema.Type.AdditionalProperties.Schema)
							sch := spec.RefProperty("#/definitions/" + nw.Name)
							lastMapSchema.NewObj = nw
							lastMapSchema.Type.AdditionalProperties.Schema = sch
							lastMapSchema.Context = lastMapSchema.Previous.Context.NewAdditionalProperty(*sch)
						}
						break
					}
					lastMapSchema.Next = &mapType{
						Previous: lastMapSchema,
						Type:     lastMapSchema.Type.AdditionalProperties.Schema,
						Context:  lastMapSchema.Context.NewAdditionalProperty(*lastMapSchema.Type.AdditionalProperties.Schema),
					}
					lastMapSchema = lastMapSchema.Next
				}
				if stack.Type != nil && stack.Type.AdditionalProperties == nil { // might not be nested, so work the simple one out here
					//b, _ := json.MarshalIndent(stack.Type, "", "  ")
					//fmt.Println(string(b))
					// pretty.Println(stack.Type)
					tpe, err := sg.TypeResolver.ResolveSchema(stack.Type, true)
					if err != nil {
						return err
					}
					// log.Printf("%s top context (complex: %t, anonymous: %t, map: %t)", sg.Name, sg.GenSchema.IsComplexObject, sg.GenSchema.IsAnonymous, sg.GenSchema.IsMap)
					// log.Printf("%s additional properties for top (complex: %t, anonymous: %t, map: %t)", tpe.GoType, tpe.IsComplexObject, tpe.IsAnonymous, tpe.IsMap)
					if tpe.IsComplexObject && tpe.IsAnonymous {
						nw := sg.makeNewStruct(sg.Name+" Anon", *stack.Type)
						sch := spec.RefProperty("#/definitions/" + nw.Name)
						stack.NewObj = nw
						stack.Type = sch
						stack.Context = sg.NewAdditionalProperty(*sch)
					}
				}

				// when additional properties, push onto the stack
				// when no more additional properties, then check for complex object
				// if it is complex and anonymous, rewrite and
				// unwind the stack so that it picks up the new ref type
				cur := lastMapSchema
				for cur != nil {
					//log.Println("popping from linked list, has previous", cur.Previous != nil)
					if cur.NewObj != nil { // this is the actual value type of a nested map
						// log.Println("adding tpe")
						if err := cur.NewObj.makeGenSchema(); err != nil {
							return err
						}
						//b, _ := json.MarshalIndent(cur.NewObj.Schema, "", "  ")
						//fmt.Println(string(b))
						sg.GenSchema.GoType += cur.NewObj.GenSchema.Name
						sg.MergeResult(cur.NewObj)
						sg.ExtraSchemas[cur.NewObj.Name] = cur.NewObj.GenSchema
					}

					if cur.Previous != nil {
						//log.Println("adding bridge map")
						if err := cur.Context.makeGenSchema(); err != nil {
							return err
						}
						//b, _ := json.MarshalIndent(cur.Previous.Type, "", "  ")
						//fmt.Println(string(b))
						if cur.NewObj != nil {
							cur.Context.MergeResult(cur.NewObj)
						} else if cur.Next != nil {
							cur.Context.MergeResult(cur.Next.Context)
						}
						cur.Previous.Context.MergeResult(cur.Context)
						cur.Previous.Context.GenSchema.AdditionalProperties = &cur.Context.GenSchema
					} else {
						if err := cur.Context.makeGenSchema(); err != nil {
							return err
						}
						sg.GenSchema.AdditionalProperties = &cur.Context.GenSchema
						if cur.NewObj != nil {
							sg.MergeResult(cur.NewObj)
						} else if cur.Next != nil {
							sg.MergeResult(cur.Next.Context)
						}

					}
					// fmt.Println("cur gotype:", cur.Context.GenSchema.GoType)
					cur = cur.Previous
				}
				sg.MergeResult(stack.Context)
				return nil
			}

			if sg.GenSchema.IsAdditionalProperties && !sg.Named {
				// log.Println("anonymous, additional complex properties")
				// for an anonoymous object, first build the new object
				// and then replace the current one with a $ref to the
				// new object
				newObj := sg.makeNewStruct(sg.GenSchema.Name+" P"+strconv.Itoa(sg.Index), sg.Schema)
				if err := newObj.makeGenSchema(); err != nil {
					return err
				}

				sg.GenSchema = GenSchema{}
				sg.Schema = *spec.RefProperty("#/definitions/" + newObj.Name)
				if err := sg.makeGenSchema(); err != nil {
					return err
				}
				sg.MergeResult(newObj)
				sg.ExtraSchemas[newObj.Name] = newObj.GenSchema
				return nil
			}
		}
	}
	return nil
}

func (sg *schemaGenContext) makeNewStruct(name string, schema spec.Schema) *schemaGenContext {
	// log.Println("making new struct:", name)
	sp := sg.TypeResolver.Doc.Spec()
	name = swag.ToGoName(name)
	if sg.TypeResolver.ModelName != sg.Name {
		name = swag.ToGoName(sg.TypeResolver.ModelName + " " + name)
	}
	sp.Definitions[name] = schema
	pg := schemaGenContext{
		Path:         "",
		Name:         name,
		Receiver:     "m",
		IndexVar:     "i",
		ValueExpr:    "m",
		Schema:       schema,
		Required:     false,
		TypeResolver: sg.TypeResolver,
		Named:        true,
		ExtraSchemas: make(map[string]GenSchema),
	}
	pg.GenSchema.IsVirtual = true

	sg.ExtraSchemas[name] = pg.GenSchema
	return &pg
}

func (sg *schemaGenContext) buildArray() error {
	// log.Println("building array")
	tpe, err := sg.TypeResolver.ResolveSchema(sg.Schema.Items.Schema, true)
	if err != nil {
		return err
	}
	// check if the element is a complex object, if so generate a new type for it
	if tpe.IsComplexObject && tpe.IsAnonymous {
		pg := sg.makeNewStruct(sg.Name+" items"+strconv.Itoa(sg.Index), *sg.Schema.Items.Schema)
		if err := pg.makeGenSchema(); err != nil {
			return err
		}
		sg.MergeResult(pg)
		sg.ExtraSchemas[pg.Name] = pg.GenSchema
		sg.Schema.Items.Schema = spec.RefProperty("#/definitions/" + pg.Name)
		if err := sg.makeGenSchema(); err != nil {
			return err
		}
		return nil
	}
	elProp := sg.NewSliceBranch(sg.Schema.Items.Schema)
	if err := elProp.makeGenSchema(); err != nil {
		return err
	}
	sg.MergeResult(elProp)
	sg.GenSchema.ItemsEnum = elProp.GenSchema.Enum
	elProp.GenSchema.Suffix = "Items"
	sg.GenSchema.GoType = "[]" + elProp.GenSchema.GoType
	sg.GenSchema.Items = &elProp.GenSchema
	return nil
}

func (sg *schemaGenContext) buildItems() error {
	presentsAsSingle := sg.Schema.Items != nil && sg.Schema.Items.Schema != nil
	if presentsAsSingle && sg.Schema.AdditionalItems != nil { // unsure if htis a valid of invalid schema
		return fmt.Errorf("single schema (%s) can't have additional items", sg.Name)
	}
	if presentsAsSingle {
		return sg.buildArray()
	}
	if sg.Schema.Items != nil {
		// This is a tuple, build a new model that represents this
		if sg.Named {
			sg.GenSchema.Name = sg.Name
			sg.GenSchema.GoType = swag.ToGoName(sg.Name)
			if sg.TypeResolver.ModelsPackage != "" {
				sg.GenSchema.GoType = sg.TypeResolver.ModelsPackage + "." + sg.GenSchema.GoType
			}
			for i, s := range sg.Schema.Items.Schemas {
				elProp := sg.NewTupleElement(&s, i)
				if err := elProp.makeGenSchema(); err != nil {
					return err
				}
				sg.MergeResult(elProp)
				elProp.GenSchema.Name = "p" + strconv.Itoa(i)
				sg.GenSchema.Properties = append(sg.GenSchema.Properties, elProp.GenSchema)
			}
			return nil
		}

		// for an anonoymous object, first build the new object
		// and then replace the current one with a $ref to the
		// new tuple object
		var sch spec.Schema
		sch.Typed("object", "")
		sch.Properties = make(map[string]spec.Schema)
		for i, v := range sg.Schema.Items.Schemas {
			sch.Required = append(sch.Required, "P"+strconv.Itoa(i))
			sch.Properties["P"+strconv.Itoa(i)] = v
		}
		sch.AdditionalItems = sg.Schema.AdditionalItems
		tup := sg.makeNewStruct(sg.GenSchema.Name+"Tuple"+strconv.Itoa(sg.Index), sch)
		if err := tup.makeGenSchema(); err != nil {
			return err
		}
		tup.GenSchema.IsTuple = true
		tup.GenSchema.IsComplexObject = false
		tup.GenSchema.Title = tup.GenSchema.Name + " a representation of an anonymous Tuple type"
		tup.GenSchema.Description = ""
		sg.ExtraSchemas[tup.Name] = tup.GenSchema

		sg.Schema = *spec.RefProperty("#/definitions/" + tup.Name)
		if err := sg.makeGenSchema(); err != nil {
			return err
		}
		sg.MergeResult(tup)

	}
	return nil
}

func (sg *schemaGenContext) buildAdditionalItems() error {
	wantsAdditionalItems :=
		sg.Schema.AdditionalItems != nil &&
			(sg.Schema.AdditionalItems.Allows || sg.Schema.AdditionalItems.Schema != nil)
	//log.Printf("%s wants additional items: %t", sg.Name, wantsAdditionalItems)

	sg.GenSchema.HasAdditionalItems = wantsAdditionalItems
	if wantsAdditionalItems {
		// check if the element is a complex object, if so generate a new type for it
		tpe, err := sg.TypeResolver.ResolveSchema(sg.Schema.AdditionalItems.Schema, true)
		if err != nil {
			return err
		}
		if tpe.IsComplexObject && tpe.IsAnonymous {
			pg := sg.makeNewStruct(sg.Name+" Items", *sg.Schema.AdditionalItems.Schema)
			if err := pg.makeGenSchema(); err != nil {
				return err
			}
			sg.Schema.AdditionalItems.Schema = spec.RefProperty("#/definitions/" + pg.Name)
			pg.GenSchema.HasValidations = true
			sg.MergeResult(pg)
			sg.ExtraSchemas[pg.Name] = pg.GenSchema
		}

		it := sg.NewAdditionalItems(sg.Schema.AdditionalItems.Schema)
		if tpe.IsInterface {
			it.Untyped = true
		}

		if err := it.makeGenSchema(); err != nil {
			return err
		}
		sg.MergeResult(it)
		sg.GenSchema.AdditionalItems = &it.GenSchema
	}
	return nil
}

func (sg *schemaGenContext) buildXMLName() error {
	if sg.Schema.XML != nil {
		//log.Printf("bulding xml name %s", sg.Name)
		sg.GenSchema.XMLName = sg.Name
		if sg.Schema.XML.Name != "" {
			sg.GenSchema.XMLName = sg.Schema.XML.Name
			if sg.Schema.XML.Attribute {
				sg.GenSchema.XMLName += ",attr"
			}
		}
	}
	return nil
}

func (sg *schemaGenContext) shortCircuitNamedRef() (bool, error) {
	// This if block ensures that a struct gets
	// rendered with the ref as embedded ref.
	if sg.Named && sg.Schema.Ref.GetURL() != nil {
		//log.Printf("short circuiting name ref: %s", sg.Schema.Ref.String())
		nullableOverride := sg.GenSchema.IsNullable
		tpe := resolvedType{}
		tpe.GoType = sg.Name
		if sg.TypeResolver.ModelsPackage != "" {
			tpe.GoType = sg.TypeResolver.ModelsPackage + "." + sg.TypeResolver.ModelName
		}

		tpe.SwaggerType = "object"
		tpe.IsComplexObject = true
		tpe.IsMap = false
		tpe.IsAnonymous = false

		item := sg.NewCompositionBranch(sg.Schema, 0)
		if err := item.makeGenSchema(); err != nil {
			return true, err
		}
		sg.GenSchema.resolvedType = tpe
		sg.GenSchema.IsNullable = sg.GenSchema.IsNullable || nullableOverride
		sg.MergeResult(item)
		sg.GenSchema.AllOf = append(sg.GenSchema.AllOf, item.GenSchema)
		return true, nil
	}
	return false, nil
}

func (sg *schemaGenContext) liftSpecialAllOf() error {
	// if there is only a $ref or a primitive and an x-isnullable schema then this is a nullable pointer
	if len(sg.Schema.AllOf) > 0 {
		var seenSchema int
		var seenNullable bool
		var schemaToLift spec.Schema

		for _, sch := range sg.Schema.AllOf {
			tpe, err := sg.TypeResolver.ResolveSchema(&sch, true)
			if err != nil {
				return err
			}
			if sg.TypeResolver.isNullable(&sch) {
				seenNullable = true
			}
			if len(sch.Type) > 0 || sch.Ref.GetURL() != nil {
				seenSchema++
				if (!tpe.IsAnonymous && tpe.IsComplexObject) || tpe.IsPrimitive {
					schemaToLift = sch
				}
			}
		}

		if seenSchema == 1 {
			//log.Printf("lifting nullable all of pattern (nullable: %t) %v", seenNullable, schemaToLift)
			sg.Schema = schemaToLift
			sg.GenSchema.IsNullable = seenNullable
		}
		return nil
	}
	return nil
}

func (sg *schemaGenContext) makeGenSchema() error {
	//log.Printf("property: (path %s) (named: %t) (name %s) (receiver %s) (indexVar %s) (expr %s) required %t", sg.Path, sg.Named, sg.Name, sg.Receiver, sg.IndexVar, sg.ValueExpr, sg.Required)
	ex := ""
	if sg.Schema.Example != nil {
		ex = fmt.Sprintf("%#v", sg.Schema.Example)
	}
	sg.GenSchema.Example = ex
	sg.GenSchema.Path = sg.Path
	sg.GenSchema.IndexVar = sg.IndexVar
	sg.GenSchema.Location = "body"
	sg.GenSchema.ValueExpression = sg.ValueExpr
	sg.GenSchema.KeyVar = sg.KeyVar
	sg.GenSchema.Name = sg.Name
	sg.GenSchema.Title = sg.Schema.Title
	sg.GenSchema.Description = sg.Schema.Description
	sg.GenSchema.ReceiverName = sg.Receiver
	sg.GenSchema.sharedValidations = sg.schemaValidations()
	sg.GenSchema.ReadOnly = sg.Schema.ReadOnly

	// b, _ := json.MarshalIndent(sg.Schema, "", "  ")
	// log.Printf("%s (named %t): %s", sg.Name, sg.Named, b)

	returns, err := sg.shortCircuitNamedRef()
	if err != nil {
		return err
	}
	if returns {
		return nil
	}
	// log.Printf("%s short circuited named ref", sg.Name)
	// b, _ = json.MarshalIndent(sg.Schema, "", "  ")
	// log.Printf("%s (named %t): %s", sg.Name, sg.Named, b)
	if err := sg.liftSpecialAllOf(); err != nil {
		return err
	}
	// log.Printf("%s short circuited all of", sg.Name)
	nullableOverride := sg.GenSchema.IsNullable
	// b, _ = json.MarshalIndent(sg.Schema, "", "  ")
	// log.Printf("%s (named %t): %s", sg.Name, sg.Named, b)

	if err := sg.buildAllOf(); err != nil {
		return err
	}
	// log.Printf("%s built all of", sg.Name)
	// b, _ = json.MarshalIndent(sg.Schema, "", "  ")
	// log.Printf("%s (named %t): %s", sg.Name, sg.Named, b)

	var tpe resolvedType
	if sg.Untyped {
		tpe, err = sg.TypeResolver.ResolveSchema(nil, !sg.Named)
	} else {
		tpe, err = sg.TypeResolver.ResolveSchema(&sg.Schema, !sg.Named)
	}
	if err != nil {
		return err
	}
	tpe.IsNullable = tpe.IsNullable || nullableOverride
	sg.GenSchema.resolvedType = tpe
	// log.Printf("%s resolved type", sg.Name)
	// b, _ = json.MarshalIndent(sg.Schema, "", "  ")
	// log.Printf("%s (named %t): %s", sg.Name, sg.Named, b)

	if err := sg.buildAdditionalProperties(); err != nil {
		return err
	}
	// b, _ = json.MarshalIndent(sg.Schema, "", "  ")
	// log.Printf("%s (named %t): %s", sg.Name, sg.Named, b)

	// log.Printf("%s built additional properties", sg.Name)
	if err := sg.buildProperties(); err != nil {
		return nil
	}
	// log.Printf("%s built properties", sg.Name)
	// b, _ = json.MarshalIndent(sg.Schema, "", "  ")
	// log.Printf("%s (named %t): %s", sg.Name, sg.Named, b)

	if err := sg.buildXMLName(); err != nil {
		return err
	}
	// log.Printf("%s built xml name", sg.Name)

	if err := sg.buildAdditionalItems(); err != nil {
		return err
	}
	// log.Printf("%s built additional items", sg.Name)

	if err := sg.buildItems(); err != nil {
		return err
	}
	// log.Printf("%s built items", sg.Name)

	return nil
}

// NOTE:
// untyped data requires a cast somehow to the inner type
// I wonder if this is still a problem after adding support for tuples
// and anonymous structs. At that point there is very little that would
// end up being cast to interface, and if it does it truly is the best guess

// GenSchema contains all the information needed to generate the code
// for a schema
type GenSchema struct {
	resolvedType
	sharedValidations
	Example                 string
	Name                    string
	Suffix                  string
	Path                    string
	ValueExpression         string
	IndexVar                string
	KeyVar                  string
	Title                   string
	Description             string
	Location                string
	ReceiverName            string
	Items                   *GenSchema
	AllowsAdditionalItems   bool
	HasAdditionalItems      bool
	AdditionalItems         *GenSchema
	Object                  *GenSchema
	XMLName                 string
	Properties              GenSchemaList
	AllOf                   []GenSchema
	HasAdditionalProperties bool
	IsAdditionalProperties  bool
	AdditionalProperties    *GenSchema
	ReadOnly                bool
	IsVirtual               bool
}

type sharedValidations struct {
	Required            bool
	MaxLength           *int64
	MinLength           *int64
	Pattern             string
	MultipleOf          *float64
	Minimum             *float64
	Maximum             *float64
	ExclusiveMinimum    bool
	ExclusiveMaximum    bool
	Enum                []interface{}
	ItemsEnum           []interface{}
	HasValidations      bool
	MinItems            *int64
	MaxItems            *int64
	UniqueItems         bool
	HasSliceValidations bool
	NeedsSize           bool
}
