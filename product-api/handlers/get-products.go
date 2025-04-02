package handlers

import (
	"net/http"
	"github.com/product-api-microservice/data"
)

// swagger:route GET /products products listProducts
// Returns a list of Products
// responses:
// 	200: productsResponse
//  500: errorResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts (w http.ResponseWriter, r *http.Request) {
	pList := data.GetProducts()

	if err := pList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}