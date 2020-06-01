// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	~ application/json
//
//	Produces:
//	~ application/json
// swagger:meta

package handlers

import "github.com/Akanibekuly/Microservices/product-api/data"

//
// NOTE: Types defined here purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defind as array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValiadtionError
}

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	//All products in the system
	// in: body
	Body []data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Create or Update.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

//swagger:parameters updateProduct
type productIDParameterWrapper struct {
	// The id of the product for which the operation relates
	//in: path
	//required: true
	ID int `json:"id"`
}
