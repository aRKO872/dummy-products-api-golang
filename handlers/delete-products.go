package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/product-api-microservice/data"
)

// swagger:route DELETE /product/{id} products deleteProduct
// Delete Product for given ID
// responses:
// 	200: noContent
//  500: errorResponse
//  404: errorResponse

// DeleteProduct 
func (*Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	delIdStr := mux.Vars(r)["id"]
	delId, convErr := strconv.Atoi(delIdStr)

	if convErr != nil {
		http.Error(w, "error decoding path parameter", http.StatusInternalServerError)
		return
	}

	err := data.DeleteProduct(delId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}