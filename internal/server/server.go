package server

import (
	"log"
	"net/http"

	"github.com/pokemonpower92/collageapi/config"
	"github.com/pokemonpower92/collageapi/internal/api"
)

type Server struct {
	conf *config.ServerConfig
	l    *log.Logger
	s    *http.Server
}

func NewServer(conf *config.ServerConfig, r *api.Router, l *log.Logger) *Server {
	ss := &http.Server{
		Addr:    conf.Host + ":" + conf.Port,
		Handler: r.Mux,
	}
	return &Server{conf, l, ss}
}

func (s *Server) Start() {
	s.l.Printf("Server starting on %s", s.conf.Host+":"+s.conf.Port)
	err := s.s.ListenAndServe()
	if err != nil {
		s.l.Fatalf("Error starting server: %s\n", err)
		panic(err)
	}
}
