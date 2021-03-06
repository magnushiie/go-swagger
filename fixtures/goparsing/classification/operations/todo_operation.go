package operations

// ListPetParams the params for the list pets query
type ListPetParams struct {
	// OutOfStock when set to true only the pets that are out of stock will be returned
	OutOfStock bool
}

// ServeAPI serves the API for this record store
func ServeAPI(host, basePath string, schemes []string) error {

	// swagger:route GET /pets pets users listPets
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	// Consumes:
	// application/json
	// application/x-protobuf
	//
	// Produces:
	// application/json
	// application/x-protobuf
	//
	// Schemes: http, https, ws, wss
	//
	// Security:
	// api_key:
	// oauth: read, write
	//
	// Responses:
	// default: genericError
	// 200: someResponse
	// 422: validationError
	mountItem("GET", basePath+"/pets", nil)

	/* swagger:route POST /pets pets users createPet

	Create a pet based on the parameters.

	Consumes:
	- application/json
	- application/x-protobuf

	Produces:
	- application/json
	- application/x-protobuf

	Schemes: http, https, ws, wss

	Responses:
	default: genericError
	200: someResponse
	422: validationError

	Security:
	api_key:
	oauth: read, write */
	mountItem("POST", basePath+"/pets", nil)

	// swagger:route GET /orders orders listOrders
	//
	// lists orders filtered by some parameters.
	//
	// Consumes:
	// application/json
	// application/x-protobuf
	//
	// Produces:
	// application/json
	// application/x-protobuf
	//
	// Schemes: http, https, ws, wss
	//
	// Security:
	// api_key:
	// oauth: read, write
	//
	// Responses:
	// default: genericError
	// 200: someResponse
	// 422: validationError
	mountItem("GET", basePath+"/orders", nil)

	// swagger:route POST /orders orders createOrder
	//
	// create an order based on the parameters.
	//
	// Consumes:
	// application/json
	// application/x-protobuf
	//
	// Produces:
	// application/json
	// application/x-protobuf
	//
	// Schemes: http, https, ws, wss
	//
	// Security:
	// api_key:
	// oauth: read, write
	//
	// Responses:
	// default: genericError
	// 200: someResponse
	// 422: validationError
	mountItem("POST", basePath+"/orders", nil)

	// swagger:route GET /orders/{id} orders orderDetails
	//
	// gets the details for an order.
	//
	// Consumes:
	// application/json
	// application/x-protobuf
	//
	// Produces:
	// application/json
	// application/x-protobuf
	//
	// Schemes: http, https, ws, wss
	//
	// Security:
	// api_key:
	// oauth: read, write
	//
	// Responses:
	// default: genericError
	// 200: someResponse
	// 422: validationError
	mountItem("GET", basePath+"/orders/:id", nil)

	// swagger:route PUT /orders/{id} orders updateOrder
	//
	// Update the details for an order.
	//
	// When the order doesn't exist this will return an error.
	//
	// Consumes:
	// application/json
	// application/x-protobuf
	//
	// Produces:
	// application/json
	// application/x-protobuf
	//
	// Schemes: http, https, ws, wss
	//
	// Security:
	// api_key:
	// oauth: read, write
	//
	// Responses:
	// default: genericError
	// 200: someResponse
	// 422: validationError
	mountItem("PUT", basePath+"/orders/:id", nil)

	// swagger:route DELETE /orders/{id} orders deleteOrder
	//
	// delete a particular order.
	//
	// Consumes:
	// application/json
	// application/x-protobuf
	//
	// Produces:
	// application/json
	// application/x-protobuf
	//
	// Schemes: http, https, ws, wss
	//
	// Security:
	// api_key:
	// oauth: read, write
	//
	// Responses:
	// default: genericError
	// 200: someResponse
	// 422: validationError
	mountItem("DELETE", basePath+"/orders/:id", nil)

	return nil
}

// not really used but I need a method to decorate the calls to
func mountItem(method, path string, handler interface{}) {}
