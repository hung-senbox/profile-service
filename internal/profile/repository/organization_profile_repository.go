package repository

import (
	"context"
	"profile-service/internal/profile/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrganizationProfileRepository interface {
	Create(ctx context.Context, organizationProfile *model.OrganizationProfile) error
	UploadSummary(ctx context.Context, organizationProfile *model.OrganizationProfile) error
	GetSummary(ctx context.Context, organizationID string) (string, error)
}

type organizationProfileRepository struct {
	collection *mongo.Collection
}

func NewOrganizationProfileRepository(collection *mongo.Collection) OrganizationProfileRepository {
	return &organizationProfileRepository{collection}
}

func (r *organizationProfileRepository) Create(ctx context.Context, organizationProfile *model.OrganizationProfile) error {
	if organizationProfile.ID.IsZero() {
		organizationProfile.ID = primitive.NewObjectID()
	}
	_, err := r.collection.InsertOne(ctx, organizationProfile)
	return err
}

func (r *organizationProfileRepository) UploadSummary(ctx context.Context, organizationProfile *model.OrganizationProfile) error {
	filter := bson.M{"organization_id": organizationProfile.OrganizationID}
	update := bson.M{
		"$set": bson.M{"summary": organizationProfile.Summary},
		"$setOnInsert": bson.M{
			"_id":             primitive.NewObjectID(),
			"organization_id": organizationProfile.OrganizationID,
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *organizationProfileRepository) GetSummary(ctx context.Context, organizationID string) (string, error) {
	filter := bson.M{"organization_id": organizationID}
	var organizationProfile model.OrganizationProfile
	err := r.collection.FindOne(ctx, filter).Decode(&organizationProfile)
	if err != nil {
		return "", err
	}
	return organizationProfile.Summary, nil
}
