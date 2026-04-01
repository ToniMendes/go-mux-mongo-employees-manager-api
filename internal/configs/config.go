// Package configs provides configuration management for the application.
package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoURI        string `envconfig:"MONGO_URI"`
	MongoDB         string `envconfig:"MONGO_DB"`
	MongoCollection string `envconfig:"MONGO_COLLECTION"`
	PORT            string `envconfig:"PORT"`
}

var Env Config

func StartConfig() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	if err := envconfig.Process("", &Env); err != nil {
		return err
	}

	return nil
}
