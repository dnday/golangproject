package main

import (
	"fmt"
	"net/http"

	"github.com/dnday/go-backend-pelatihan-kmteti/src/handler"
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
	h.HandleFunc("/api/employeelist", handler.EmployeeListHandler)
	// h.HandleFunc("api/book", getBook)
	// h.HandleFunc("api/books", listBookHandler)
	// h.HandleFunc("api/book/add", addBook)
	// h.HandleFunc("api/book/update", updateBook)
	// h.HandleFunc("api/book/delete", deleteBook)
	// h.HandleFunc("api/employees", getEmployees)
	// h.HandleFunc("api/employee/add", addEmployee)

	fmt.Println("HTTP Server running on port 8080")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
