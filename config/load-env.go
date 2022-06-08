package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	return port
}
