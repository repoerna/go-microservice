package handlers

import (
	"log"
	"net/http"
	"time"
)

type Goodbye struct {
	log *log.Logger
}

func NewGoodbye(log *log.Logger) *Goodbye {
	return &Goodbye{log: log}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 10)

	w.Write([]byte("goodbye!!"))
}
