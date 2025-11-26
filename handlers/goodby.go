package handlers

import (
	"log"
	"net/http"
)

type Goodby struct {
	l *log.Logger
}

func NewGoodBy(l *log.Logger) *Goodby {
	return &Goodby{l}
}

func (g *Goodby) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeeee"))
}
