package data

import "testing"

func TestProductValidation(t *testing.T) {
	p := &Product{
		Name:  "tets",
		Price: 1.00,
		SKU:   "qwe-qwer-qwer",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
