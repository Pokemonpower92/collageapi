package api

import (
	"log"
	"net/http"
)

type ImageSetHandler struct {
	l *log.Logger
}

func NewImageSetHandler(l *log.Logger) *ImageSetHandler {
	return &ImageSetHandler{l}
}

func (ish *ImageSetHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ish.l.Println("Handle ImageSet")
	if r.Method == http.MethodPost {
		ish.CreateImageSet(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (ish *ImageSetHandler) CreateImageSet(rw http.ResponseWriter, r *http.Request) {
	ish.l.Println("Handle POST ImageSet")
}
