package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func GetProducts() Products {
	return products
}

var products = Products{
	{
		ID:          1,
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       1.99,
		SKU:         "ABC123",
		CreatedAt:   "2020-01-01 00:00:00",
		UpdatedAt:   "2020-01-01 00:00:00",
		DeletedAt:   "2020-01-01 00:00:00",
	},
	{
		ID:          2,
		Name:        "Product 2",
		Description: "Product 2 description",
		Price:       2.99,
		SKU:         "ABC124",
		CreatedAt:   "2020-01-01 00:00:00",
		UpdatedAt:   "2020-01-01 00:00:00",
		DeletedAt:   "2020-01-01 00:00:00",
	},
}
