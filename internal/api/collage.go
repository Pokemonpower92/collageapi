package api

import (
	"log"
	"net/http"
)

type CollageHandler struct {
	l *log.Logger
}

func NewCollageHandler(l *log.Logger) *CollageHandler {
	return &CollageHandler{l}
}

func (ch *CollageHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ch.l.Println("Handle Collage")
	if r.Method == http.MethodPost {
		ch.CreateCollage(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (ch *CollageHandler) CreateCollage(rw http.ResponseWriter, r *http.Request) {
	ch.l.Println("Handle POST Collage")
}
