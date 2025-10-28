package repository

import (
	"context"
	"fmt"
	"profile-service/internal/profile/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OwnerCodeRepository interface {
	Create(ctx context.Context, ownerCode *model.OwnerCode) error
	FindByOwnerIDAndRole(ctx context.Context, ownerID string, ownerRole string) (*model.OwnerCode, error)
}

type ownerCodeRepository struct {
	collection *mongo.Collection
}

func NewOwnerCodeRepository(collection *mongo.Collection) OwnerCodeRepository {
	return &ownerCodeRepository{collection}
}

func (r *ownerCodeRepository) Create(ctx context.Context, ownerCode *model.OwnerCode) error {
	// Kiểm tra trùng owner_id + owner_role
	filter := bson.M{
		"owner_id":   ownerCode.OwnerID,
		"owner_role": ownerCode.OwnerRole,
	}
	existing := r.collection.FindOne(ctx, filter)
	if existing.Err() == nil {
		return fmt.Errorf("owner code for owner_id=%s and role=%s already exists", ownerCode.OwnerID, ownerCode.OwnerRole)
	}
	if existing.Err() != mongo.ErrNoDocuments {
		return existing.Err()
	}

	// Tiếp tục tạo
	if ownerCode.ID.IsZero() {
		ownerCode.ID = primitive.NewObjectID()
	}
	ownerCode.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, ownerCode)
	return err
}

func (r *ownerCodeRepository) FindByOwnerIDAndRole(ctx context.Context, ownerID string, ownerRole string) (*model.OwnerCode, error) {
	filter := bson.M{
		"owner_id":   ownerID,
		"owner_role": ownerRole,
	}

	var ownerCode model.OwnerCode
	err := r.collection.FindOne(ctx, filter).Decode(&ownerCode)
	if err != nil {
		return nil, err
	}

	return &ownerCode, nil
}
