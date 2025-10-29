package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StaffProfile struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	StaffID   string             `json:"staff_id" bson:"staff_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Email     string             `json:"email" bson:"email"`
	Phone     string             `json:"phone" bson:"phone"`
}
