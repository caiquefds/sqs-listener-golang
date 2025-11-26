package config

import (
	"log"

	"github.com/joho/godotenv"
)

type EnvConfig struct{}

func (e EnvConfig) configure() {
	log.Printf("Loanding env...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
