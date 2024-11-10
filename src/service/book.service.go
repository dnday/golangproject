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
	ID     string  `json:"id" bson:"id"`
	Title  string  `json:"title" bson:"title"`
	Author string  `json:"author" bson:"author"`
	Stock  int     `json:"stock" bson:"stock"`
	Price  float64 `json:"price" bson:"price"`
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
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Stock  int     `json:"stock"`
	Price  float64 `json:"price"`
}

func GetAllBook() (*BookResponse, error) {
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

	var booksList []*Book

	for cur.Next(context.TODO()) {
		var books model.Book
		cur.Decode(&books)
		booksList = append(booksList, &Book{
			ID:     books.ID.Hex(),
			Title:  books.Title,
			Author: books.Author,
			Stock:  books.Stock,
			Price:  books.Price,
		})
	}
	return &BookResponse{
		Data: booksList,
	}, nil
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

func GetBookDetail() (*BookResponse, error) {
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

	var booksList []*Book

	for cur.Next(context.TODO()) {
		var books model.Book
		cur.Decode(&books)
		booksList = append(booksList, &Book{
			ID:     books.ID.Hex(),
			Title:  books.Title,
			Author: books.Author,
			Stock:  books.Stock,
			Price:  books.Price,
		})
	}
	return &BookResponse{
		Data: booksList,
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

	coll := db.MongoDB.Collection("product")
	_, err = coll.InsertOne(context.TODO(), model.Book{
		ID:     primitive.NewObjectID(),
		Title:  bookReq.Title,
		Author: bookReq.Author,
		Stock:  bookReq.Stock,
		Price:  bookReq.Price,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}

	return nil
}
