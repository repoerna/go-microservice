package handlers

import (
	"go-microservice/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		p.log.Println("POST /products")
		p.addProducts(w, r)
		return
	}

	// /products/{id}
	if r.Method == http.MethodPut {

		// p.log.Println("PUT /products")
		// expect id on the URI
		path := r.URL.Path

		// p.log.Println(path)

		re := regexp.MustCompile(`/([0-9]+)`)

		g := re.FindAllStringSubmatch(path, -1)

		// p.log.Println(g)

		if len(g) != 1 || g == nil {
			http.Error(w, "Invalid URI 1", http.StatusBadRequest)
		}

		if len(g[0]) != 2 {
			http.Error(w, "Invalid URI 2", http.StatusBadRequest)
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.log.Println("Invalid URI unable to convert to number", idString)
			http.Error(w, "Invalid URI 3", http.StatusBadRequest)
			return
		}

		p.updateProduct(w, r, id)

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET products")

	products := data.GetProducts()

	err := products.ToJson(w)

	if err != nil {
		http.Error(w, "unable marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(w http.ResponseWriter, r *http.Request) {
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

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request, id int) {
	p.log.Println("Handle PUT products")

	product := &data.Product{}

	err := product.FromJson(r.Body)

	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "unable to update product", http.StatusInternalServerError)
	}
}
