package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when loading env file: %s", err.Error())
	}
	return &Config{
		Database: Database{
			DBUser: os.Getenv("DBUSER"),
			DBPass: os.Getenv("DBPASS"),
			DBHost: os.Getenv("DBHOST"),
			DBPort: os.Getenv("DBPORT"),
			DBName: os.Getenv("DBNAME"),
		},
		Midtrans: Midtrans{
			Key:    os.Getenv("MIDTRANS_KEY"),
			IsProd: os.Getenv("MIDTRANS_ENV") == "production",
		},
	}
}