package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Citizen struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	DOB       string             `json:"dob"`
	Gender    string             `json:"gender"`
	Address   string             `json:"address"`
	City      string             `json:"city"`
	State     string             `json:"state"`
	Pincode   string             `json:"pincode"`
}
