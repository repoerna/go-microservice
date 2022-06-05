package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.log.Println("Hello from HandelFunc")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error occured"))

		return
	}

	fmt.Fprintf(w, "hello %s\n", d)

}
