package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dnday/go-backend-pelatihan-kmteti/src/service"
)

func OneBookHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := service.GetBookDetail()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
