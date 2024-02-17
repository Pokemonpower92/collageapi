package main

import (
	"log"
	"os"

	"github.com/pokemonpower92/collageapi/config"
	"github.com/pokemonpower92/collageapi/internal/api"
	"github.com/pokemonpower92/collageapi/internal/server"
)

func main() {
	l := log.New(os.Stdout, "collageapi", log.LstdFlags)
	c := config.NewServerConfig()
	r := api.NewRouter()
	s := server.NewServer(c, r, l)

	s.Start()
}
