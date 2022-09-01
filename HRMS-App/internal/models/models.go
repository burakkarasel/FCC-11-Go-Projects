package models

import "go.mongodb.org/mongo-driver/mongo"

// Employee holds the employee data
type Employee struct {
	Name   string  `json:"name"`
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// MongoInstance holds our DB connection
type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}
