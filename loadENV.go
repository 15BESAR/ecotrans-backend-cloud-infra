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
	appName   string
	sigKeyJwt []byte
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
		appName:   os.Getenv("APPLICATION_NAME"),
		sigKeyJwt: []byte(os.Getenv("JWT_SIGNATURE_KEY"))}
	return environment
}
