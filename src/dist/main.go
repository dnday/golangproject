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

	h.HandleFunc("/api/listbooks", handler.ListBookHandler)
	h.HandleFunc("/api/onebook", handler.OneBookHandler)
	h.HandleFunc("/api/updatebook", handler.UpdateBookHandler)
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

// func listBookHandler(w http.ResponseWriter, r *http.Request) {
// 	collection := client.Database("bookstore").Collection("books")
// 	cursor, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(context.TODO())

// 	var books []Book
// 	if err = cursor.All(context.TODO(), &books); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(books)
// }

// func getBook(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	collection := client.Database("bookstore").Collection("books")
// 	var book Book
// 	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(book)
// }

// func addBook(w http.ResponseWriter, r *http.Request) {
// 	var book Book
// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	collection := client.Database("bookstore").Collection("books")
// 	_, err = collection.InsertOne(context.TODO(), book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	var book Book
// 	err := json.NewDecoder(r.Body).Decode(&book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	coll := client.Database("bookstore").Collection("books")
// 	update := bson.M{
// 		"$set": bson.M{
// 			"stock": book.Stock,
// 			"price": book.Price,
// 		},
// 	}
// 	_, err = coll.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	collection := client.Database("bookstore").Collection("books")
// 	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func getEmployees(w http.ResponseWriter, r *http.Request) {
// 	collection := client.Database("bookstore").Collection("employees")
// 	cursor, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer cursor.Close(context.TODO())

// 	var employees []Employee
// 	if err = cursor.All(context.TODO(), &employees); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(employees)
// }

// func addEmployee(w http.ResponseWriter, r *http.Request) {
// 	var employee Employee
// 	err := json.NewDecoder(r.Body).Decode(&employee)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	collection := client.Database("bookstore").Collection("employees")
// 	_, err = collection.InsertOne(context.TODO(), employee)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// }
