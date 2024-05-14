package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("error when loading env file: %s, Environment from Docker Will Be Load", err.Error())
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
		DigitalOceanSpaces: DigitalOceanSpaces{
			AccessToken: os.Getenv("DO_SPACES_ACCESS_TOKEN"),
			SecretKey:   os.Getenv("DO_SPACES_SECRET_KEY"),
			Region:      os.Getenv("DO_SPACES_REGION"),
			Name:        os.Getenv("DO_SPACES_NAME"),
			Endpoint:    os.Getenv("DO_SPACES_ENDPOINT"),
		},
	}
}
