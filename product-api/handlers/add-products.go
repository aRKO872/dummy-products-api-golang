package handlers

import (
	// "encoding/json"
	"net/http"

	// "regexp"
	// "strconv"

	"github.com/product-api-microservice/data"
)

// swagger:route POST /products products createProducts
// Create new products and add to list of products
// responses:
// 	200: productsResponse
//  500: errorResponse

func (p *Products) AddProduct (w http.ResponseWriter, r *http.Request) {
	inputProduct := r.Context().Value(data.ProductKey).(data.Product)

	w.Header().Set("Content-Type", "application/json")

	pList := data.AddProduct(inputProduct)

	if err := pList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}