package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID    string  `json:"id" bson:"_id,omitempty"`
	Title string  `json:"title" bson:"title"`
	Stock int     `json:"stock" bson:"stock"`
	Price float64 `json:"price" bson:"price"`
}

var client *mongo.Client

func init() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
}

func EditBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database("mydb").Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": bookID}
	update := bson.M{
		"$set": bson.M{
			"stock": updatedBook.Stock,
			"price": updatedBook.Price,
		},
	}

	var book Book
	err = collection.FindOneAndUpdate(ctx, filter, update).Decode(&book)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
