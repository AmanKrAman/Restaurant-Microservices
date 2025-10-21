package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DBUrl: os.Getenv("DB_URL"),
		Port:  os.Getenv("PORT"),
	}

	if cfg.DBUrl == "" {
		log.Fatal("DB_URL not set")
	}
	if cfg.Port == "" {
		cfg.Port = "8083"
	}

	return cfg
}
