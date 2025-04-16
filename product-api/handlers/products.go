// Package Classification of Product API
//
// # Documentation for Product API
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

	"github.com/aRKO872/currency-grpc-service/protos/currency"
	"github.com/product-api-microservice/data"
)

type Products struct {
	l *log.Logger
	cc currency.CurrencyClient
}

func NewProducts (l *log.Logger, cc *currency.CurrencyClient) *Products {
	return &Products{
		l: l,
		cc: *cc,
	}
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