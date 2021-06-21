package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DC_AUTH string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	cfg := &Config{
		DC_AUTH: os.Getenv("DC_AUTH"),
	}
	log.Info("Config loaded succesfully")

	return cfg

}
