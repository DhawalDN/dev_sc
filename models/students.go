package models

import (
	"context"
	"dev_sc/db"
	"dev_sc/forms"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Gender   string             `json:"gender"`
	IsActive bool               `json:"isActive"`
}

type StudentModel struct{}

var server = "127.0.0.1"

var client = db.GetClient()

func (m *StudentModel) Create(data forms.CreateStudentCommand) error {

	collection := client.Database("db-mgo").Collection("students")

	insertResult, err := collection.InsertOne(context.TODO(), bson.M{"name": data.Name, "email": data.Email, "password": data.Password, "gender": data.Gender, "isActive": true})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return err
}

func (m *StudentModel) Find() (list []*Student, err error) {

	collection := client.Database("db-mgo").Collection("students")
	cur, err := collection.Find(context.TODO(), bson.M{})
	for cur.Next(context.TODO()) {
		var li Student
		err = cur.Decode(&li)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		list = append(list, &li)
	}

	return list, err
}

func (m *StudentModel) Get(id string) (student Student, err error) {
	collection := client.Database("db-mgo").Collection("students")
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	err = collection.FindOne(context.TODO(), filter).Decode(&student)

	return student, err
}

func (m *StudentModel) Update(id string, data forms.CreateStudentCommand) (err error) {

	return err
}

func (m *StudentModel) Delete(id string) (err error) {

	return err
}
