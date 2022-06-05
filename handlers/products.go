package handlers

import (
	"go-microservice/data"
	"log"
	"net/http"
)

type Products struct {
	log *log.Logger
}

func NewProduct(log *log.Logger) *Products {
	return &Products{log}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.log.Println("GET /products")
		p.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.log.Println("Hello from product handler")

	products := data.GetProducts()

	err := products.ToJson(w)

	if err != nil {
		http.Error(w, "unable marshal json", http.StatusInternalServerError)
	}
}
