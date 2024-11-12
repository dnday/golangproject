package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	NIK               int                `bson:"nik"`
	TertieryEducation string             `bson:"tertieryeducation"`
	EntryDate         string             `bson:"entrydate"`
	Status            string             `bson:"status"`
}
