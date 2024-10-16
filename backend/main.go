// main.go - Initializes Go API server, PostgreSQL connection, and Redis caching.
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Chrisadams777/AI-Code-Search-Tool/backend/api"
	"github.com/Chrisadams777/AI-Code-Search-Tool/backend/cache"
	"github.com/Chrisadams777/AI-Code-Search-Tool/backend/config"
	"github.com/Chrisadams777/AI-Code-Search-Tool/backend/db"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server...")

	// Load environment variables
	config.LoadEnv()

	// Initialize database connection
	db.InitDB()

	// Initialize Redis cache
	cache.InitRedis()

	// Set up router and routes
	router := mux.NewRouter()
	router.HandleFunc("/search/keyword", api.KeywordSearch).Methods("GET")
	router.HandleFunc("/search/regex", api.RegexSearch).Methods("GET")
	router.HandleFunc("/search/semantic", api.SemanticSearch).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

