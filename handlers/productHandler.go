package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "product-management-example/helpers"
	. "product-management-example/models"
	"strconv"
	"time"
)

var productStore = make(map[string]Product)
var id int = 0

// HTTP POST - /api/products
func PostProductsHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)

	CheckError(err)

	product.CreatedOn = time.Now()
	id++

	key := strconv.Itoa(id)
	product.ID = key
	productStore[key] = product

	data, err := json.Marshal(product)

	CheckError(err)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// HTTP GET - /api/products/{id}
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product

	for _, product := range productStore {
		products = append(products, product)
	}

	data, err := json.Marshal(products)

	CheckError(err)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP GET - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key := vars["id"]

	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)

	CheckError(err)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP PUT - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)

	key := vars["id"]

	var prodUpt Product

	err = json.NewDecoder(r.Body).Decode(&prodUpt)

	CheckError(err)

	if _, ok := productStore[key]; ok {
		prodUpt.ID = key
		prodUpt.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodUpt
	} else {
		log.Printf("Deger bulunamadi : %s", key)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	msj := "updated"
	w.Write([]byte(msj))

}

// HTTP DELETE - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["id"]

	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Deger bulunamadi : %s", key)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	msj := "deleted"
	w.Write([]byte(msj))
}
