package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Author    string             `bson:"author"`
	PrintYear int                `bson:"printyear"`
	Stock     int                `bson:"stock"`
	Price     float64            `bson:"price"`
}
