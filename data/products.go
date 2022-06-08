package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return products
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	products = append(products, p)
}

func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = id
	products[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (int, error) {

	for i, p := range products {
		log.Println(p)
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrProductNotFound
}

func getNextID() int {
	lastProductID := products[len(products)-1].ID
	return lastProductID + 1
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
