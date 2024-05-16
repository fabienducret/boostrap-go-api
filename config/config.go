package config

import "os"

type Config struct {
	Port string
}

const DEFAULT_PORT = "8080"

func MustInit() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	return &Config{
		Port: port,
	}
}
