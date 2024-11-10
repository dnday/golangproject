package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dnday/go-backend-pelatihan-kmteti/src/service"
	"github.com/gorilla/mux"
)

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var bookReq service.BookRequest
	if err := json.NewDecoder(r.Body).Decode(&bookReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	bookReqBytes, err := json.Marshal(bookReq)
	if err != nil {
		http.Error(w, "Failed to encode request payload", http.StatusInternalServerError)
		return
	}

	err = service.UpdateBook(strconv.Itoa(bookID), bytes.NewReader(bookReqBytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book updated successfully"))
}
