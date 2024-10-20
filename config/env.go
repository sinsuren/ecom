package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress: fmt.Sprintf("%s:%s",
			getEnv("DB_ADDRESS", "tadf"),
			getEnv("DB_PORT", "fdafad"),
		),
		DBName: getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
