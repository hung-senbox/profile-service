package repository

import (
	"context"
	"profile-service/internal/profile/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OwnerCodeRepository interface {
	Create(ctx context.Context, ownerCode *model.OwnerCode) error
}

type ownerCodeRepository struct {
	collection *mongo.Collection
}

func NewOwnerCodeRepository(collection *mongo.Collection) OwnerCodeRepository {
	return &ownerCodeRepository{collection}
}

func (r *ownerCodeRepository) Create(ctx context.Context, ownerCode *model.OwnerCode) error {
	if ownerCode.ID.IsZero() {
		ownerCode.ID = primitive.NewObjectID()
	}
	ownerCode.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, ownerCode)
	return err
}
