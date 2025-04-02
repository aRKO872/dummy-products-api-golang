package handlers

import "github.com/product-api-microservice/data"

// A list of Products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIdParameterWrapper struct {
	// The id of product to delete from DB
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type productNoContentWrapper struct {
}

// swagger:parameters updateProducts createProducts
type productParamsWrapper struct {
	// Product data structure for creating or updating products

	// required: true
	// in: body
	Body data.Product
}

// swagger:response errorResponse
type errorResponseWrapper struct {
	// Generic error for endpoints

	// in: body
	Body GenericError
}

// swagger:response validationError
type validationErrorResponseWrapper struct {
	// Collection of errors for failed validation

	// in: body
	Body ValidationError
}

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}
