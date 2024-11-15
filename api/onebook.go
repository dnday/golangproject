package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dnday/go-backend-pelatihan-kmteti/src/service"
)

func OneBookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bookID := r.URL.Query().Get("id")
		bookTitle := r.URL.Query().Get("title")

		if bookID == "" && bookTitle == "" {
			http.Error(w, "missing book id or title", http.StatusBadRequest)
			return
		}

		var book interface{}
		var err error

		if bookID != "" {
			book, err = service.GetBookByID(bookID)
		} else {
			book, err = service.GetBookByTitle(bookTitle)
		}

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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
		return
	case "POST":
	default:
		log.Default().Println(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
