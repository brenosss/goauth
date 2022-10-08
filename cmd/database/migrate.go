package main

import (
	db "backend/pkg/database"
)

func main() {
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
