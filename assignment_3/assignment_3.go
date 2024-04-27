package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strconv"
)

const (
	host     = "localhost"
	user     = "postgres"
	port     = 5432
	password = "19082004"
	dbname   = "postgres"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func getProductHandler(db *sql.DB, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get product ID from URL
		vars := mux.Vars(r)
		productID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		// check Redis cache for data
		ctx := r.Context()
		cachedData, err := rdb.Get(ctx, "product:"+vars["id"]).Result()
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cachedData))
			return
		}

		// product not found in cache
		var product Product
		row := db.QueryRow("SELECT id, name, description, price FROM products WHERE id = $1", productID)
		err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err == sql.ErrNoRows {
			// NOT FOUND
			http.NotFound(w, r)
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// cache data in Redis
		jsonData, _ := json.Marshal(product)
		rdb.Set(ctx, "product:"+vars["id"], jsonData, 0)

		// Return data to client with 200
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func main() {
	pq := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pq)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-13672.c300.eu-central-1-1.ec2.redns.redis-cloud.com:13672",
		Password: "eq9mWyPcf7yy1U0PkGQjZbA7t3Q5hR4c",
		DB:       0,
	})

	// router to work with HTTP and URL's
	r := mux.NewRouter()

	r.HandleFunc("/products/{id}", getProductHandler(db, rdb)).Methods("GET")

	// Start HTTP server
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
