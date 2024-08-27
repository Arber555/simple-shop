package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func main() {
	http.HandleFunc("/products", productsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Products API server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := fetchProductsFromDB()
	if err != nil {
		log.Printf("Failed to fetch products: %v", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func fetchProductsFromDB() ([]Product, error) {
	connStr := os.Getenv("PRODUCTS_DB_DSN")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
