package repository

import (
	"context"
	"profile-service/internal/profile/model"

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
	_, err := r.collection.InsertOne(ctx, ownerCode)
	return err
}
