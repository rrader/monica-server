package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type errorResponse struct {
	Status string
	Error  string
}

// GETBarcodeHandler handles http requests querying the barcode by it's code
func GETBarcodeHandler(w http.ResponseWriter, r *http.Request) {
	db := GetDB()
	vars := mux.Vars(r)

	var barcode Barcode
	err := db.QueryRow("SELECT id, code, name FROM barcode WHERE code=$1", vars["code"]).Scan(&barcode.ID, &barcode.Code, &barcode.Name)

	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if err == sql.ErrNoRows {
		w.Header().Set("Content-Type", "application/json")
		response := errorResponse{
			Status: "ERROR",
			Error:  "Not found",
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	// response := BarcodeResponse(barcode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barcode)
}
