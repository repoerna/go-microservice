// Package classification for Product API
//
// Documentation for Product API
//
// 	Scheme: http
// 	Basepath: /products
// 	Version: 1.0.0
//
// 	Consumes:
// 	- aplication/json
//
// 	Produce:
// 	- application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"go-microservice/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	log *log.Logger
}

func NewProduct(log *log.Logger) *Products {
	return &Products{log}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET products")

	products := data.GetProducts()

	err := products.ToJson(w)

	if err != nil {
		http.Error(w, "unable marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST products")

	product := &data.Product{}
	err := product.FromJson(r.Body)

	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		return
	}

	p.log.Printf("Adding product: %#v", product)
	data.AddProduct(product)

}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle PUT products")

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	p.log.Println("Updating product:", id)

	p.log.Println(r.Context())

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, product)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "unable to update product", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		err := product.FromJson(r.Body)

		if err != nil {
			p.log.Println("error deserializing product:", err)
			http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validatthe product
		err = product.Validate()
		if err != nil {
			p.log.Println("error validating product:", err)
			http.Error(w, fmt.Sprintf("error validating product: %s", err.Error()), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})

}
