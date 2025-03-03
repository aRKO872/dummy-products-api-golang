package handlers

import (
	// "encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/product-api-microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts (l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		path := r.URL.Path

		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(path, -1)

		prodId, err := strconv.Atoi(g[0][1])

		if err != nil {
			http.Error(w, "error getting path param", http.StatusInternalServerError)
			return
		}

		p.UpdateProduct(w, r, prodId)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) UpdateProduct (
	w http.ResponseWriter, 
	r *http.Request,
	prodId int,
) {
	updateReq := new(data.Product)

	defer r.Body.Close()

	if err := updateReq.FromBody(r.Body); err != nil {
		http.Error(w, "error reading from body and unmarshalling request", http.StatusInternalServerError)
		return
	}

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

	prodList[resProdInd] = *updateReq
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
	inputProduct := new(data.Product)
	
	defer r.Body.Close()

	if err := inputProduct.FromBody(r.Body); err != nil {
		http.Error(w, "error unmarshalling product", http.StatusInternalServerError)
		return
	}

	pList := data.GetProducts()

	ind := len(pList)+1
	inputProduct.ID = ind

	pList = append(pList, *inputProduct)

	data.SetProducts(pList)

	if err := pList.ToJSON(w); err != nil {
		http.Error(w, "error marshalling product list", http.StatusInternalServerError)
		return
	}
}