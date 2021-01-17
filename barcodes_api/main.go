package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, world!")

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/barcode/{code:\\d*}", GETBarcodeHandler)
	http.ListenAndServe(":8080", r)
}
