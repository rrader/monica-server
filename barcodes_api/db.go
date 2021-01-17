package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Barcode entity from DataBase
type Barcode struct {
	ID   int
	Code string
	Name string
}

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "monica_storage"
)

var db *sql.DB

// OpenConnection opens connection to the db
func openConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

// GetDB returns database connection singleton
func GetDB() *sql.DB {
	return db
}

func init() {
	db = openConnection()
}
