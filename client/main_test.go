package main

import (
	"fmt"
	"product-api-microservice-client/client"
	"product-api-microservice-client/client/products"
	"testing"
)
func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8000")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	prodRaw, err := c.Products.ListProducts(products.NewListProductsParams())

	if err != nil {
		t.Fatal(err)
	}

	prods := prodRaw.Payload
	
	for _, prod := range prods {
		fmt.Println(prod.Name)
	}
}