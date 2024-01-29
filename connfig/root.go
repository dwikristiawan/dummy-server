package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

type Root struct {
	Server Server
}

func Load(filenames ...string) *Root {
	err := godotenv.Overload(filenames...)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	r := Root{
		Server: Server{},
	}
	mustLoad("SERVER", &r.Server)
	return &r
}

func mustLoad(prefix string, spec interface{}) {
	err := envconfig.Process(prefix, spec)
	if err != nil {
		panic(err)
	}
}
