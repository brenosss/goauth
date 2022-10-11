package database

import (
	"backend/pkg/config"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const migrationDirName = "migrations"

//go:embed migrations/*
var embedMigrations embed.FS

const queriesDirName = "queries"

//go:embed queries/*
var embedQueries embed.FS

func GetQuery(queryName string) string {
	sqlFile, err := embedQueries.ReadFile(queriesDirName + "/" + queryName)
	if err != nil {
		panic(err)
	}
	return string(sqlFile)
}

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
		_, err = conn.Exec(string(sqlFile))
		if err != nil {
			panic(err)
		}
	}
	conn.Close()
}

func GetConnection() *sql.DB {
	dbPath := "/sqlite/" + config.GetConfig().DatabaseName
	db, err := sql.Open("sqlite3", config.GetProjectDir()+dbPath)
	if err != nil {
		panic(err)
	}
	return db
}

func ClearDatabase() {
	if config.GetConfig().Env != "test" {
		panic("You can only clear the database in test environment")
	}
	dbPath := "/sqlite/" + config.GetConfig().DatabaseName
	err := os.Remove(config.GetProjectDir() + dbPath)
	if err != nil {
		panic(err)
	}
}
