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
)

type EmployeeRequest struct {
	Name              string `json:"name"`
	NIK               int    `json:"nik"`
	TertieryEducation string `json:"tertieryeducation"`
	EntryDate         string `json:"entrydate"`
	Status            string `json:"status"`
}

type EmployeeSummaries struct {
	Name      string `json:"name" bson:"name"`
	EntryDate string `json:"entrydate" bson:"entrydate"`
	Status    string `json:"status" bson:"status"`
}

type EmployeeResponseSummaries struct {
	Data []*EmployeeSummaries `json:"data"`
}

func GetEmployeeDetail() (*EmployeeResponseSummaries, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("employee")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var employeeListSummaries []*EmployeeSummaries

	for cur.Next(context.TODO()) {
		var employee model.Employee
		cur.Decode(&employee)
		employeeListSummaries = append(employeeListSummaries, &EmployeeSummaries{
			Name:      employee.Name,
			EntryDate: employee.EntryDate,
			Status:    employee.Status,
		})
	}
	return &EmployeeResponseSummaries{Data: employeeListSummaries},
		nil
}

func AddEmployee(req io.Reader) error {
	var emplReq EmployeeRequest
	err := json.NewDecoder(req).Decode(&emplReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("employee")
	_, err = coll.InsertOne(context.TODO(), model.Employee{
		Name:              emplReq.Name,
		NIK:               emplReq.NIK,
		TertieryEducation: emplReq.TertieryEducation,
		EntryDate:         emplReq.EntryDate,
		Status:            emplReq.Status,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}

	return nil
}
