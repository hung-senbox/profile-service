package repository

import (
	"context"
	"profile-service/internal/profile/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentInformationRepository interface {
	Create(ctx context.Context, studentInformation *model.StudentInformation) error
	Update(ctx context.Context, studentInformation *model.StudentInformation) error
	UploadStudentInfo(ctx context.Context, studentInformation *model.StudentInformation) error
	GetByStudentID(ctx context.Context, studentID string) (*model.StudentInformation, error)
}

type studentInformationRepository struct {
	collection *mongo.Collection
}

func NewStudentInformationRepository(collection *mongo.Collection) StudentInformationRepository {
	return &studentInformationRepository{collection}
}

func (r *studentInformationRepository) Create(ctx context.Context, studentInformation *model.StudentInformation) error {
	if studentInformation.ID.IsZero() {
		studentInformation.ID = primitive.NewObjectID()
	}
	_, err := r.collection.InsertOne(ctx, studentInformation)
	return err
}

func (r *studentInformationRepository) Update(ctx context.Context, studentInformation *model.StudentInformation) error {
	filter := bson.M{"_id": studentInformation.ID}
	update := bson.M{
		"$set": studentInformation,
	}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *studentInformationRepository) Delete(ctx context.Context, studentInformation *model.StudentInformation) error {

	filter := bson.M{"_id": studentInformation.ID}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}

func (r *studentInformationRepository) GetByStudentID(ctx context.Context, studentID string) (*model.StudentInformation, error) {
	filter := bson.M{"student_id": studentID}
	var studentInformation model.StudentInformation
	err := r.collection.FindOne(ctx, filter).Decode(&studentInformation)
	if err != nil {
		return nil, err
	}
	return &studentInformation, nil
}

func (r *studentInformationRepository) UploadStudentInfo(ctx context.Context, studentInformation *model.StudentInformation) error {
	filter := bson.M{"student_id": studentInformation.StudentID}

	// Prepare update with $set for all fields and $setOnInsert for ID
	update := bson.M{
		"$set": bson.M{
			"student_id":           studentInformation.StudentID,
			"dob":                  studentInformation.DOB,
			"gender":               studentInformation.Gender,
			"study_level":          studentInformation.StudyLevel,
			"min_water_must_drink": studentInformation.MinWaterMustDrink,
			"description":          studentInformation.Description,
			"mode":                 studentInformation.Mode,
		},
		"$setOnInsert": bson.M{
			"_id": primitive.NewObjectID(),
		},
	}
	_, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}
