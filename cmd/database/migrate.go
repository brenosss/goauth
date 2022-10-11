package main

import (
	"backend/pkg/config"
	db "backend/pkg/database"
)

func main() {
	config.LoadEnv("dev")
	db.ApplyMigrations()
}
