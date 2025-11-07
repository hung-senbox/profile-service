package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrganizationProfile struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	OrganizationID string             `json:"organization_id" bson:"organization_id"`
	Summary        string             `json:"summary" bson:"summary"`
}
