package data

import (
	"testing"
)

func TestSKUValidatorFunc(t *testing.T) {
	prod := Product{
		Name: "fdjuhdfiog",
		Price: 6,
		SKU: "sdg567",
	}

	if err := prod.Validate(); err != nil {
		t.Fatal(err)
	}
}
