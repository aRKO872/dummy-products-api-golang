package handlers

import (
	"net/http"
	"github.com/product-api-microservice/data"
)

// swagger:route PUT /products products updateProducts
// Update Product from the list and returns an updated list of products
// responses:
// 	200: productsResponse
//  404: errorResponse
//  500: errorResponse
//  400: validationError

func (p *Products) UpdateProduct (
	w http.ResponseWriter, 
	r *http.Request,
) {
	updateReq := r.Context().Value(data.ProductKey).(data.Product)

	w.Header().Set("Content-Type", "application/json")

	prodList, err := data.UpdateProduct(updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := prodList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}