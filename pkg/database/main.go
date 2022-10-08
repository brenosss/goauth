package database

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*
var embedMigrations embed.FS

const migrationDirName = "migrations"

func GetMigrations() []fs.DirEntry {
	fileMigrations, err := embedMigrations.ReadDir(migrationDirName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Available migrations:")
	for _, file := range fileMigrations {
		fmt.Println(file.Name())
	}
	return fileMigrations
}

func ApplyMigrations() {
	fileMigrations := GetMigrations()
	conn := GetConnection()
	for _, file := range fileMigrations {
		fmt.Println("Applying migration:", file.Name())
		sqlFile, err := embedMigrations.ReadFile(migrationDirName + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		sqlStatements := strings.Split(string(sqlFile), "\n")
		for _, sqlStatement := range sqlStatements {
			fmt.Println("SQL:", string(sqlStatement))
			_, err := conn.Exec(string(sqlFile))
			if err != nil {
				panic(err)
			}
		}
	}
	conn.Close()
}

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	return db
}
