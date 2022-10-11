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
	Env          string
}

const projectDirName = "tags"

func LoadEnv(env string) {
	err := godotenv.Load(GetProjectDir() + "/.env-" + env)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetProjectDir() string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}

func GetConfig() Config {

	return Config{
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Env:          os.Getenv("ENV"),
	}
}
