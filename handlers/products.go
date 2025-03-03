package handlers

import (
	// "encoding/json"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "regexp"
	// "strconv"

	"github.com/gorilla/mux"
	"github.com/product-api-microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts (l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) UpdateProduct (
	w http.ResponseWriter, 
	r *http.Request,
) {
	vars := mux.Vars(r)

	prodId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "error decoding path parameter", http.StatusInternalServerError)
		return
	}

	// updateReq := new(data.Product)

	// defer r.Body.Close()

	// if err := updateReq.FromBody(r.Body); err != nil {
	// 	http.Error(w, "error reading from body and unmarshalling request", http.StatusInternalServerError)
	// 	return
	// }

	updateReq := r.Context().Value(data.ProductKey).(data.Product)

	prodList := data.GetProducts()

	resProdInd := -1

	for prodInd, p := range prodList {
		if p.ID == prodId {
			resProdInd = prodInd
			break
		}
	}

	if resProdInd == -1 {
		http.Error(w, "product not found!", http.StatusNotFound)
		return
	}

	prodList[resProdInd] = updateReq
	data.SetProducts(prodList)

	if err := prodList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}

func (p *Products) GetProducts (w http.ResponseWriter, r *http.Request) {
	pList := data.GetProducts()

	if err := pList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct (w http.ResponseWriter, r *http.Request) {
	// inputProduct := new(data.Product)
	
	// defer r.Body.Close()

	// if err := inputProduct.FromBody(r.Body); err != nil {
	// 	http.Error(w, "error unmarshalling product", http.StatusInternalServerError)
	// 	return
	// }

	inputProduct := r.Context().Value(data.ProductKey).(data.Product)

	pList := data.GetProducts()

	ind := len(pList)+1
	inputProduct.ID = ind

	pList = append(pList, inputProduct)

	data.SetProducts(pList)

	if err := pList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
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
			http.Error(w, fmt.Sprintf("error during validation of request : %s", err.Error()), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), data.ProductKey, *inputProduct)
		updatedReq := r.WithContext(ctx)

		next.ServeHTTP(w, updatedReq)
	})
} 