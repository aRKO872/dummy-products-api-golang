// Package Classification of Product API
//
// Documentation for Product API
// 
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	//	"encoding/json"
	"context"
	"fmt"
	"log"
	"net/http"

	//	"regexp"
	//	"strconv"

	"github.com/product-api-microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts (l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) MiddlewareEncodingProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inputProduct := new(data.Product)
	
		defer r.Body.Close()

		if err := inputProduct.FromBody(r.Body); err != nil {
			http.Error(w, "error unmarshalling product", http.StatusInternalServerError)
			return
		}

		if err := inputProduct.Validate(); err != nil {
			http.Error(w, fmt.Sprintf("error during validation of request : %s", err.Error()), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), data.ProductKey, *inputProduct)
		updatedReq := r.WithContext(ctx)

		next.ServeHTTP(w, updatedReq)
	})
} 