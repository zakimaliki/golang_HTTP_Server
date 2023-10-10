package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var products []Product

func main() {
	http.HandleFunc("/product", ProductHandler)
	http.ListenAndServe(":8080", nil)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Handle GET request, return list of products
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	} else if r.Method == "POST" {
		// Handle POST request, parse request body and validate
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid request body")
			return
		}

		// Basic validation
		if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid product data")
			return
		}

		// Save the product to the list
		products = append(products, product)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}