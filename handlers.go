package api

import (
	"encoding/json"
	"net/http"
	"backend/search"
	//"backend/utils"
	"time"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	start := time.Now()
	results := search.Search(keyword)
	duration := time.Since(start)

	response := map[string]interface{}{
		"results":  results,
		"count":    len(results),
		"duration": duration.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}