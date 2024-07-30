package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublisHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		PublisHost:   GetEnv("PUBLIC_HOST", "localhost"),
		Port:         GetEnv("PORT", "8080"),
		DBUser:       GetEnv("DB_USER", "root"),
		DBPassword:   GetEnv("DB_PASSWORD", ""),
		DBName:       GetEnv("DB_NAME", "ecom"),
		DBAddress:    fmt.Sprintf("%s:%s", GetEnv("DB_HOST", "localhost"), GetEnv("DB_PORT", "3307")),
	}
}

func GetEnv(key, fallback string)string{
	if value, ok := os.LookupEnv(key);ok{
		return value
	}
	return fallback
}