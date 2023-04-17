package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envs struct {
	APP_ENV string
}

func makeEnvs() *envs {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured loading env file. Err: %s", err)
	}

	envs := &envs{
		APP_ENV: os.Getenv("APP_ENV"),
	}

	return envs
}

var Envs = makeEnvs()
