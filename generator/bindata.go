// Code generated by go-bindata.
// sources:
// templates/client/client.gotmpl
// templates/client/facade.gotmpl
// templates/client/parameter.gotmpl
// templates/client/response.gotmpl
// templates/docstring.gotmpl
// templates/model.gotmpl
// templates/modelvalidator.gotmpl
// templates/schemabody.gotmpl
// templates/schematype.gotmpl
// templates/server/builder.gotmpl
// templates/server/configureapi.gotmpl
// templates/server/main.gotmpl
// templates/server/operation.gotmpl
// templates/server/parameter.gotmpl
// templates/structfield.gotmpl
// templates/validation/customformat.gotmpl
// templates/validation/primitive.gotmpl
// templates/validation/structfield.gotmpl
// DO NOT EDIT!

package generator

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// templatesClientClientGotmpl reads file data from disk. It returns an error on failure.
func templatesClientClientGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/client/client.gotmpl"
	name := "templates/client/client.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesClientFacadeGotmpl reads file data from disk. It returns an error on failure.
func templatesClientFacadeGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/client/facade.gotmpl"
	name := "templates/client/facade.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesClientParameterGotmpl reads file data from disk. It returns an error on failure.
func templatesClientParameterGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/client/parameter.gotmpl"
	name := "templates/client/parameter.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesClientResponseGotmpl reads file data from disk. It returns an error on failure.
func templatesClientResponseGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/client/response.gotmpl"
	name := "templates/client/response.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesDocstringGotmpl reads file data from disk. It returns an error on failure.
func templatesDocstringGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/docstring.gotmpl"
	name := "templates/docstring.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesModelGotmpl reads file data from disk. It returns an error on failure.
func templatesModelGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/model.gotmpl"
	name := "templates/model.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesModelvalidatorGotmpl reads file data from disk. It returns an error on failure.
func templatesModelvalidatorGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/modelvalidator.gotmpl"
	name := "templates/modelvalidator.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesSchemabodyGotmpl reads file data from disk. It returns an error on failure.
func templatesSchemabodyGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/schemabody.gotmpl"
	name := "templates/schemabody.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesSchematypeGotmpl reads file data from disk. It returns an error on failure.
func templatesSchematypeGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/schematype.gotmpl"
	name := "templates/schematype.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesServerBuilderGotmpl reads file data from disk. It returns an error on failure.
func templatesServerBuilderGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/server/builder.gotmpl"
	name := "templates/server/builder.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesServerConfigureapiGotmpl reads file data from disk. It returns an error on failure.
func templatesServerConfigureapiGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/server/configureapi.gotmpl"
	name := "templates/server/configureapi.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesServerMainGotmpl reads file data from disk. It returns an error on failure.
func templatesServerMainGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/server/main.gotmpl"
	name := "templates/server/main.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesServerOperationGotmpl reads file data from disk. It returns an error on failure.
func templatesServerOperationGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/server/operation.gotmpl"
	name := "templates/server/operation.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesServerParameterGotmpl reads file data from disk. It returns an error on failure.
func templatesServerParameterGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/server/parameter.gotmpl"
	name := "templates/server/parameter.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesStructfieldGotmpl reads file data from disk. It returns an error on failure.
func templatesStructfieldGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/structfield.gotmpl"
	name := "templates/structfield.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesValidationCustomformatGotmpl reads file data from disk. It returns an error on failure.
func templatesValidationCustomformatGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/validation/customformat.gotmpl"
	name := "templates/validation/customformat.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesValidationPrimitiveGotmpl reads file data from disk. It returns an error on failure.
func templatesValidationPrimitiveGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/validation/primitive.gotmpl"
	name := "templates/validation/primitive.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// templatesValidationStructfieldGotmpl reads file data from disk. It returns an error on failure.
func templatesValidationStructfieldGotmpl() (*asset, error) {
	path := "/home/ivan/go/src/github.com/go-swagger/go-swagger/generator/templates/validation/structfield.gotmpl"
	name := "templates/validation/structfield.gotmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/client/client.gotmpl": templatesClientClientGotmpl,
	"templates/client/facade.gotmpl": templatesClientFacadeGotmpl,
	"templates/client/parameter.gotmpl": templatesClientParameterGotmpl,
	"templates/client/response.gotmpl": templatesClientResponseGotmpl,
	"templates/docstring.gotmpl": templatesDocstringGotmpl,
	"templates/model.gotmpl": templatesModelGotmpl,
	"templates/modelvalidator.gotmpl": templatesModelvalidatorGotmpl,
	"templates/schemabody.gotmpl": templatesSchemabodyGotmpl,
	"templates/schematype.gotmpl": templatesSchematypeGotmpl,
	"templates/server/builder.gotmpl": templatesServerBuilderGotmpl,
	"templates/server/configureapi.gotmpl": templatesServerConfigureapiGotmpl,
	"templates/server/main.gotmpl": templatesServerMainGotmpl,
	"templates/server/operation.gotmpl": templatesServerOperationGotmpl,
	"templates/server/parameter.gotmpl": templatesServerParameterGotmpl,
	"templates/structfield.gotmpl": templatesStructfieldGotmpl,
	"templates/validation/customformat.gotmpl": templatesValidationCustomformatGotmpl,
	"templates/validation/primitive.gotmpl": templatesValidationPrimitiveGotmpl,
	"templates/validation/structfield.gotmpl": templatesValidationStructfieldGotmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"client": &bintree{nil, map[string]*bintree{
			"client.gotmpl": &bintree{templatesClientClientGotmpl, map[string]*bintree{
			}},
			"facade.gotmpl": &bintree{templatesClientFacadeGotmpl, map[string]*bintree{
			}},
			"parameter.gotmpl": &bintree{templatesClientParameterGotmpl, map[string]*bintree{
			}},
			"response.gotmpl": &bintree{templatesClientResponseGotmpl, map[string]*bintree{
			}},
		}},
		"docstring.gotmpl": &bintree{templatesDocstringGotmpl, map[string]*bintree{
		}},
		"model.gotmpl": &bintree{templatesModelGotmpl, map[string]*bintree{
		}},
		"modelvalidator.gotmpl": &bintree{templatesModelvalidatorGotmpl, map[string]*bintree{
		}},
		"schemabody.gotmpl": &bintree{templatesSchemabodyGotmpl, map[string]*bintree{
		}},
		"schematype.gotmpl": &bintree{templatesSchematypeGotmpl, map[string]*bintree{
		}},
		"server": &bintree{nil, map[string]*bintree{
			"builder.gotmpl": &bintree{templatesServerBuilderGotmpl, map[string]*bintree{
			}},
			"configureapi.gotmpl": &bintree{templatesServerConfigureapiGotmpl, map[string]*bintree{
			}},
			"main.gotmpl": &bintree{templatesServerMainGotmpl, map[string]*bintree{
			}},
			"operation.gotmpl": &bintree{templatesServerOperationGotmpl, map[string]*bintree{
			}},
			"parameter.gotmpl": &bintree{templatesServerParameterGotmpl, map[string]*bintree{
			}},
		}},
		"structfield.gotmpl": &bintree{templatesStructfieldGotmpl, map[string]*bintree{
		}},
		"validation": &bintree{nil, map[string]*bintree{
			"customformat.gotmpl": &bintree{templatesValidationCustomformatGotmpl, map[string]*bintree{
			}},
			"primitive.gotmpl": &bintree{templatesValidationPrimitiveGotmpl, map[string]*bintree{
			}},
			"structfield.gotmpl": &bintree{templatesValidationStructfieldGotmpl, map[string]*bintree{
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

