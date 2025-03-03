package data

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,gt=5"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"deletedOn,omitempty"`
}

type Products []Product

func (p *Product) FromBody(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func SKUValidatorFunc (fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-z]{3}[0-9]{3}$`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return (len(matches) == 1) 
}

func (p *Product) Validate() error {
	vldt := validator.New()
	vldt.RegisterValidation("sku", SKUValidatorFunc)

	return vldt.Struct(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func SetProducts(prods Products) {
	productList = prods
}

var productList = []Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       29.99,
		SKU:         "anc234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk",
		Price:       29.99,
		SKU:         "ack234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
