// main.go - Initializes Go API server, PostgreSQL connection, and Redis caching.
package main

import (
	"fmt"
	"log"
	"net/http"
	"search-app/api"
	"search-app/cache"
	"search-app/config"
	"search-app/db"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv() // Loads environment variables (e.g., DB, API keys)
	db.InitDB()      // Connects to PostgreSQL
	cache.InitRedis() // Connects to Redis

	router := mux.NewRouter()
	router.HandleFunc("/search/keyword", api.KeywordSearch).Methods("GET")
	router.HandleFunc("/search/regex", api.RegexSearch).Methods("GET")
	router.HandleFunc("/search/semantic", api.SemanticSearch).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// search.go - Implements Sourcegraph search requests.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var sourcegraphAPI = os.Getenv("SOURCEGRAPH_API_URL")

func KeywordSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	resp, err := http.Get(fmt.Sprintf("%s/search?q=%s&type=code", sourcegraphAPI, query))
	if err != nil {
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	json.NewEncoder(w).Encode(result)
}

