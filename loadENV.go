package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	port      string
	apiKey    string
	dbUser    string
	dbPass    string
	dbName    string
	dbTCPHost string
	dbPort    string
	appName   string
	sigKeyJwt []byte
	projectID string
}

var env Env

func LoadEnvironment() Env {
	godotenv.Load()
	environment := Env{
		port:      os.Getenv("PORT"),
		apiKey:    os.Getenv("API_KEY"),
		dbUser:    os.Getenv("DB_USERNAME"),
		dbName:    os.Getenv("DB_NAME"),
		dbPass:    os.Getenv("DB_PASSWORD"),
		dbTCPHost: os.Getenv("DB_TCP_HOST"),
		dbPort:    os.Getenv("DB_PORT"),
		appName:   os.Getenv("APPLICATION_NAME"),
		sigKeyJwt: []byte(os.Getenv("JWT_SIGNATURE_KEY")),
		projectID: os.Getenv("PROJECT_ID"),
	}
	return environment
}
