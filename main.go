package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from HandelFunc")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error occured"))

			return
		}

		fmt.Fprintf(w, "hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye")
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Error")

		http.Error(w, "this is error endpoint", http.StatusBadRequest)
	})

	http.ListenAndServe(":9090", nil)
}
