package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentInformation struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	StudentID         string             `json:"student_id" bson:"student_id"`
	DOB               time.Time          `json:"dob" bson:"dob"`
	Gender            string             `json:"gender" bson:"gender"`
	StudyLevel        uint               `json:"study_level" bson:"study_level"`
	MinWaterMustDrink uint               `json:"min_water_must_drink" bson:"min_water_must_drink"`
	Description       string             `json:"description" bson:"description"`
	Mode              string             `json:"mode" bson:"mode"`
}
