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
