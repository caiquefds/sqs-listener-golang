package config

import "log"

func Configure() {
	log.Println("Configure application settings...")
	EnvConfig{}.configure()
	AWSConfig{}.configure()
	log.Println("Application configured!")
}
