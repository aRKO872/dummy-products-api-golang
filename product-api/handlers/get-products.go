package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	w.Header().Set("Content-Type", "application/json")

	if err := pList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}

func (p *Products) GetProductsSingle (w http.ResponseWriter, r *http.Request) {
	mp := mux.Vars(r)
	id, err := strconv.Atoi(mp["id"])
	if err != nil {
		http.Error(w, "not a valid ID", http.StatusBadRequest)
		return
	}

	pList := data.GetProducts()
	prodInd := -1
	for ind, prod := range pList {
		if prod.ID == id {
			prodInd = ind
			break
		}
	}

	if prodInd == -1 {
		http.Error(w, "product doesn't exist", http.StatusNotFound)
		return
	}

	finalProd := pList[prodInd]

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	if err := enc.Encode(finalProd); err != nil {
		http.Error(w, "error marshalling product", http.StatusInternalServerError)
		return
	}
}