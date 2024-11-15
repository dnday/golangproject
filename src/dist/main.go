package main

import (
	"fmt"
	"net/http"

	handler "github.com/dnday/go-backend-pelatihan-kmteti/api"
)

func main() {
	h := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	h.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Server!!"))
	})

	h.HandleFunc("/api/books", handler.ListBookHandler)
	h.HandleFunc("/api/onebook", handler.OneBookHandler)
	h.HandleFunc("/api/employees", handler.EmployeeListHandler)

	fmt.Println("HTTP Server running on port 8080")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
