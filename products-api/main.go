package main

import (
	"database/sql"
	"encoding/json"
	"io"
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

type App struct {
	DB *sql.DB
}

func main() {
	// Set up logging
	logFile, err := os.OpenFile("/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Log to both file and standard output
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	connStr := os.Getenv("PRODUCTS_DB_DSN")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	app := &App{DB: db}

	http.HandleFunc("/products", app.productsHandler)
	http.HandleFunc("/health", app.healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Products API server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (app *App) productsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := app.fetchProductsFromDB()
	if err != nil {
		log.Printf("Failed to fetch products: %v", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (app *App) fetchProductsFromDB() ([]Product, error) {
	rows, err := app.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Println("Failed to execute query:", err)
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

func (app *App) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.DB.Ping(); err != nil {
		log.Println("Health check failed: database unreachable")
		http.Error(w, "Database unreachable", http.StatusInternalServerError)
		return
	}

	log.Println("Health check passed")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
