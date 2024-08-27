package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Stock int    `json:"stock"`
}

type Stock struct {
	ProductID int `json:"product_id"`
	Stock     int `json:"stock"`
}

func main() {
	fmt.Println("Starting web", os.Getenv("PRODUCTS_API_URL"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Welcome to the Product Store!</h1>")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := fetchProducts()
		if err != nil {
			log.Printf("Failed to fetch products: %v", err)
			http.Error(w, "Unable to fetch products", http.StatusInternalServerError)
			return
		}

		stock, err := fetchStock()
		if err != nil {
			log.Printf("Failed to fetch stock: %v", err)
			http.Error(w, "Unable to fetch stock", http.StatusInternalServerError)
			return
		}

		stockMap := make(map[int]int)
		for _, s := range stock {
			stockMap[s.ProductID] = s.Stock
		}

		for i := range products {
			products[i].Stock = stockMap[products[i].ID]
		}

		json.NewEncoder(w).Encode(products)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting web server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func fetchProducts() ([]Product, error) {
	productsURL := os.Getenv("PRODUCTS_API_URL")
	productsURL, err := url.JoinPath(productsURL, "/products")
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(productsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	return products, err
}

func fetchStock() ([]Stock, error) {
	stockURL := os.Getenv("STOCK_API_URL")
	stockURL, err := url.JoinPath(stockURL, "/stock")
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(stockURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var stock []Stock
	err = json.NewDecoder(resp.Body).Decode(&stock)
	return stock, err
}
