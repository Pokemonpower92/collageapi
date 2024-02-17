package config

import "os"

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: os.Getenv("HOST_NAME"),
		Port: os.Getenv("PORT"),
	}
}
