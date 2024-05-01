package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

var db *gorm.DB

func InitDB(config Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("error when connecting to the database : %s", err.Error())
	}
	log.Println("connected to the database")
	Migrate()
	return db
}

func Migrate() {
	err := db.AutoMigrate()
	if err != nil {
		log.Fatalf("error migratin database: %s", err.Error())
	}
}

func InitConfigMySQL() Config {
	return Config{
		DBUser: os.Getenv("DBUSER"),
		DBPass: os.Getenv("DBPASS"),
		DBHost: os.Getenv("DBHOST"),
		DBPort: os.Getenv("DBPORT"),
		DBName: os.Getenv("DBNAME"),
	}
}
