package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Config struct to hold configuration from the file
type Config struct {
	Name string `json:"name"`
	Port string `json:"port"`
}

// Stock represents the stock information for a product
type Stock struct {
	ProductID int `json:"product_id"`
	Stock     int `json:"stock"`
}

var config Config

func main() {
	// Load the configuration from the file
	err := loadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	http.HandleFunc("/stock", stockHandler)

	// Use the configured port if available, otherwise fallback to env PORT or default to 8080
	port := config.Port
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
	}
	log.Printf("%s server started on port %s", config.Name, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// stockHandler handles the /stock endpoint
func stockHandler(w http.ResponseWriter, r *http.Request) {
	stock := []Stock{
		{ProductID: 1, Stock: 50},
		{ProductID: 2, Stock: 30},
		{ProductID: 3, Stock: 75},
	}

	json.NewEncoder(w).Encode(stock)
}

// loadConfig loads configuration from the given filepath
func loadConfig(filepath string) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Parse the JSON configuration into the config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	return nil
}
