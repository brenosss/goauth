package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type Config struct {
	// Database configuration
	DatabaseName string
}

const projectDirName = "tags"

func LoadEnv(env string) {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + "/.env-" + env)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() Config {

	return Config{
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}
}
