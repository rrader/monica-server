package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

// Barcode entity from DataBase
type Barcode struct {
	ID   int `json:"-"`
	Code string
	Name string
}

var (
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

func getenvInt(key string, fallback int) int {
	var value string
	value = os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	valInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return valInt
}

func getenvString(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func init() {
	dbHost = getenvString("MONICA_DBHOST", dbHost)
	dbPort = getenvInt("MONICA_DBPORT", dbPort)
	dbUser = getenvString("MONICA_DBUSER", dbUser)
	dbPassword = getenvString("MONICA_DBPASS", dbPassword)
	dbName = getenvString("MONICA_DBNAME", dbName)

	db = openConnection()
}
