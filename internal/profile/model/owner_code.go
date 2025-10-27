package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OwnerCode struct {
	ID           primitive.ObjectID `bson:"_id"`
	OwnerID      string             `bson:"owner_id"`
	OwnerRole    string             `bson:"owner_role"`
	Code         string             `bson:"code"`
	CreatedIndex int                `bson:"created_index"`
	CreatedAt    time.Time          `bson:"created_at"`
}
