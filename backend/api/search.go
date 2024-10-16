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

func RegexSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	resp, err := http.Get(fmt.Sprintf("%s/search?q=%s&type=regexp", sourcegraphAPI, query))
	if err != nil {
		http.Error(w, "Regex search failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func SemanticSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	resp, err := http.Get(fmt.Sprintf("%s/search?q=%s&type=semantic", sourcegraphAPI, query))
	if err != nil {
		http.Error(w, "Semantic search failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	json.NewEncoder(w).Encode(result)
}
