package api

import (
	"log"
	"net/http"
	"os"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	sm := http.NewServeMux()
	l := log.New(os.Stdout, "api", log.LstdFlags)
	sm.Handle("/collage", NewCollageHandler(l))
	sm.Handle("/imageset", NewImageSetHandler(l))

	return &Router{Mux: sm}
}
