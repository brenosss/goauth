package main

import (
	db "backend/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ApplyMigrations()
	/*
		conn.Exec("INSERT INTO users (name) VALUES ('breno')")
		rows, err := conn.Query("SELECT * FROM users")
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			i := User{}
			err = rows.Scan(&i.Id, &i.Username)
			if err != nil {
				panic(err)
			}
			println(i.Id, i.Username)
		}
	*/
}
