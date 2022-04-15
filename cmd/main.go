package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "product-management-example/handlers"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductHandler).Methods("GET")

	r.HandleFunc("/api/products", PostProductsHandler).Methods("POST")

	r.HandleFunc("/api/products/{id}", PutProductHandler).Methods("PUT")

	r.HandleFunc("/api/products/{id}", DeleteProductHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	server.ListenAndServe()
	log.Println("Server ending...")
}
