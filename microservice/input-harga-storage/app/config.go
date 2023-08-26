package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser string
	DBPass string
	DBName string
	DBPort string
	DBHost string
}

func InitConfig(f string) *DBConfig {
	err := godotenv.Load(f)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := &DBConfig{
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
		DBHost: os.Getenv("DB_HOST"),
	}

	return conf
}
