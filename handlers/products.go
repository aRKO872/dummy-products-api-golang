package handlers

import (
	// "encoding/json"
	"log"
	"net/http"

	"github.com/product-api-microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts (l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()

	// prodStr, err := json.Marshal(productList)

	// if err != nil {
	// 	http.Error(w, "error marshalling product list", http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(prodStr)

	if err := productList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}