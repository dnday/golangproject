package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dnday/go-backend-pelatihan-kmteti/src/service"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := service.AddBook(r.Body)
	if err != nil {
		if err.Error() == "bad request" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("book added successfully")
}
