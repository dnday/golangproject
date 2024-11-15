package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dnday/go-backend-pelatihan-kmteti/src/service"
)

func ListBookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data, err := service.GetBookSummaries()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	case "POST":
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
		return
	case "PUT":
		bookID := r.URL.Query().Get("id")
		if bookID == "" {
			http.Error(w, "missing book id", http.StatusBadRequest)
			return
		}
		err := service.UpdateBook(bookID, r.Body)
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
		json.NewEncoder(w).Encode("book updated successfully")
		return
	case "DELETE":
		bookID := r.URL.Query().Get("id")
		if bookID == "" {
			http.Error(w, "missing book id", http.StatusBadRequest)
			return
		}
		err := service.DeleteBook(bookID)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("book deleted successfully")
		return
	default:
		log.Default().Println(http.StatusMethodNotAllowed)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
