package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, loading environment variables from system")
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}

func GetMongoURI() string {
	return mustGetEnv("MONGO_URI")
}

func GetMongoDBName() string {
	return mustGetEnv("MONGO_DB_NAME")
}

func GetPort() string {
	return getEnv("PORT", "8080")
}
