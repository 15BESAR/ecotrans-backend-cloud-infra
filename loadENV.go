package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	port   string
	apiKey string
	dbUser string
	dbPass string
}

func LoadEnvironment() Env {
	godotenv.Load()
	environment := Env{port: os.Getenv("PORT"), apiKey: os.Getenv("API_KEY"), dbUser: os.Getenv("DB_USERNAME"), dbPass: os.Getenv("DB_PASSWORD")}
	return environment
}
