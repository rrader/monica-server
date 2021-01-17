package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GETBarcodeHandler handles http requests querying the barcode by it's code
func GETBarcodeHandler(w http.ResponseWriter, r *http.Request) {
	// db := GetDB()
	vars := mux.Vars(r)
	println(vars["code"])
}
