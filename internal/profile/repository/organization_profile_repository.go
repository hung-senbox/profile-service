package repository

import (
	"context"
	"profile-service/internal/profile/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrganizationProfileRepository interface {
	Create(ctx context.Context, organizationProfile *model.OrganizationProfile) error
	UploadSummary(ctx context.Context, organizationProfile *model.OrganizationProfile) error
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
	update := bson.M{"$set": bson.M{"summary": organizationProfile.Summary}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
