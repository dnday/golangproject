package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/dnday/go-backend-pelatihan-kmteti/src/db"
	"github.com/dnday/go-backend-pelatihan-kmteti/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        string  `json:"id" bson:"_id"`
	Title     string  `json:"title" bson:"title"`
	Author    string  `json:"author" bson:"author"`
	PrintYear int     `json:"printyear" bson:"printyear"`
	Stock     int     `json:"stock" bson:"stock"`
	Price     float64 `json:"price" bson:"price"`
}

type BookResponse struct {
	Data []*Book `json:"data"`
}

type BookSummary struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type BookSummaryResponse struct {
	Data []*BookSummary `json:"data"`
}

type BookRequest struct {
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	PrintYear int     `json:"printyear"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
}

func GetBookSummaries() (*BookSummaryResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var bookSummaries []*BookSummary

	for cur.Next(context.TODO()) {
		var books model.Book
		cur.Decode(&books)
		bookSummaries = append(bookSummaries, &BookSummary{
			Title:  books.Title,
			Author: books.Author,
			Price:  books.Price,
		})
	}
	return &BookSummaryResponse{
		Data: bookSummaries,
	}, nil
}

func GetBookByID(bookID string) (*Book, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, errors.New("invalid book ID")
	}

	var book model.Book
	err = coll.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("book not found")
	}

	return &Book{
		Title:     book.Title,
		Author:    book.Author,
		PrintYear: book.PrintYear,
		Price:     book.Price,
		Stock:     book.Stock,
	}, nil
}

func GetBookByTitle(bookTitle string) (*Book, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	objTitle := bookTitle

	var book model.Book
	err = coll.FindOne(context.TODO(), bson.M{"title": objTitle}).Decode(&book)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("book not found")
	}

	return &Book{
		Title:     book.Title,
		Author:    book.Author,
		PrintYear: book.PrintYear,
		Price:     book.Price,
		Stock:     book.Stock,
	}, nil
}
func AddBook(req io.Reader) error {
	var bookReq BookRequest
	err := json.NewDecoder(req).Decode(&bookReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	_, err = coll.InsertOne(context.TODO(), model.Book{
		Title:     bookReq.Title,
		Author:    bookReq.Author,
		PrintYear: bookReq.PrintYear,
		Stock:     bookReq.Stock,
		Price:     bookReq.Price,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func UpdateBook(bookID string, bookData io.Reader) error {
	var bookReq BookRequest
	err := json.NewDecoder(bookData).Decode(&bookReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return errors.New("invalid book ID")
	}

	update := bson.M{
		"$set": bson.M{
			"stock": bookReq.Stock,
			"price": bookReq.Price,
		},
	}

	_, err = coll.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}

	return nil
}

func DeleteBook(bookID string) error {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return errors.New("invalid book ID")
	}

	_, err = coll.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}

	return nil
}
