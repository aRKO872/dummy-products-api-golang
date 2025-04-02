package data

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure of an API product
// swagger:model Product
type Product struct {
	// the id of the product
	// min: 1
	// required: true
	ID          int     `json:"id"`

	// name of the product
	// required: true
	Name        string  `json:"name" validate:"required"`

	// description of product
	Description string  `json:"description"`

	// Price of Product
	// min: 5
	// required: true
	Price       float32 `json:"price" validate:"required,gt=5"`

	// SKU of Product 
	// pattern: ^[a-z]{3}[0-9]{3}$
	// example: abc123
	// required: true
	SKU         string  `json:"sku" validate:"required,sku"`
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

func AddProduct(prod Product) Products {
	ind := len(productList)+1
	prod.ID = ind
	productList = append(productList, prod)

	return productList
}

func UpdateProduct(prod Product) (Products, error) {
	resProdInd := -1

	for prodInd, p := range productList {
		if p.ID == prod.ID {
			resProdInd = prodInd
			break
		}
	}

	if resProdInd == -1 {
		return Products{}, errors.New("product not found!")
	}

	productList[resProdInd] = prod

	return productList, nil
}

func DeleteProduct(prodId int) (error) {
	foundFlag := 0

	for ind := 0; ind < len(productList); ind += 1 {
		if productList[ind].ID == prodId {
			foundFlag = 1
			productList = append(productList[:ind], productList[ind+1:]...)
			ind -= 1
		}
	}

	if foundFlag == 0 {
		return errors.New("product not found")
	}

	return nil
}

var productList = []Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       29.99,
		SKU:         "anc234",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk",
		Price:       29.99,
		SKU:         "ack234",
	},
}