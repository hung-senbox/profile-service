package usecase

import (
	"context"
	"profile-service/helper"
	"profile-service/internal/cache/caching"
	"profile-service/internal/profile/model"
	"profile-service/internal/profile/repository"
	"profile-service/pkg/constants"
)

type GenerateOwnerCodeUseCase interface {
	GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error)
	GenerateTeacherCode(ctx context.Context, teacherID string, createdIndex int) (string, error)
	GenerateStaffCode(ctx context.Context, staffID string, createdIndex int) (string, error)
	GenerateParentCode(ctx context.Context, parentID string, createdIndex int) (string, error)
	GenerateUserCode(ctx context.Context, userID string, createdIndex int) (string, error)
	GenerateChildCode(ctx context.Context, userID string, createdIndex int) (string, error)
}

type generateOwnerCodeUseCase struct {
	ownerCodeRepo  repository.OwnerCodeRepository
	cachingService caching.CachingService
}

func NewGenerateOwnerCodeUseCase(ownerCodeRepo repository.OwnerCodeRepository, cachingService caching.CachingService) GenerateOwnerCodeUseCase {
	return &generateOwnerCodeUseCase{
		ownerCodeRepo:  ownerCodeRepo,
		cachingService: cachingService,
	}
}

func (s *generateOwnerCodeUseCase) GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error) {
	code := helper.GenerateStudentCode(createdIndex)

	ownerCode := &model.OwnerCode{
		OwnerID:      studentID,
		OwnerRole:    string(constants.StudentRole),
		Code:         code,
		CreatedIndex: createdIndex,
	}

	err := s.ownerCodeRepo.Create(ctx, ownerCode)
	if err != nil {
		return "", err
	}

	_ = s.cachingService.SetStudentCode(ctx, studentID, code)
	return code, nil
}

func (s *generateOwnerCodeUseCase) GenerateTeacherCode(ctx context.Context, teacherID string, createdIndex int) (string, error) {
	code := helper.GenerateTeacherCode(createdIndex)

	ownerCode := &model.OwnerCode{
		OwnerID:      teacherID,
		OwnerRole:    string(constants.TeacherRole),
		Code:         code,
		CreatedIndex: createdIndex,
	}

	err := s.ownerCodeRepo.Create(ctx, ownerCode)
	if err != nil {
		return "", err
	}

	_ = s.cachingService.SetTeacherCode(ctx, teacherID, code)
	return code, nil
}

func (s *generateOwnerCodeUseCase) GenerateStaffCode(ctx context.Context, staffID string, createdIndex int) (string, error) {
	code := helper.GenerateStaffCode(createdIndex)

	ownerCode := &model.OwnerCode{
		OwnerID:      staffID,
		OwnerRole:    string(constants.StaffRole),
		Code:         code,
		CreatedIndex: createdIndex,
	}

	err := s.ownerCodeRepo.Create(ctx, ownerCode)
	if err != nil {
		return "", err
	}

	_ = s.cachingService.SetStaffCode(ctx, staffID, code)
	return code, nil
}

func (s *generateOwnerCodeUseCase) GenerateParentCode(ctx context.Context, parentID string, createdIndex int) (string, error) {
	code := helper.GenerateParentCode(createdIndex)

	ownerCode := &model.OwnerCode{
		OwnerID:      parentID,
		OwnerRole:    string(constants.ParentRole),
		Code:         code,
		CreatedIndex: createdIndex,
	}

	err := s.ownerCodeRepo.Create(ctx, ownerCode)
	if err != nil {
		return "", err
	}

	_ = s.cachingService.SetParentCode(ctx, parentID, code)
	return code, nil
}

func (s *generateOwnerCodeUseCase) GenerateUserCode(ctx context.Context, userID string, createdIndex int) (string, error) {
	code := helper.GenerateUserCode(createdIndex)

	ownerCode := &model.OwnerCode{
		OwnerID:      userID,
		OwnerRole:    string(constants.UserRole),
		Code:         code,
		CreatedIndex: createdIndex,
	}

	err := s.ownerCodeRepo.Create(ctx, ownerCode)
	if err != nil {
		return "", err
	}

	_ = s.cachingService.SetUserCode(ctx, userID, code)
	return code, nil
}

func (s *generateOwnerCodeUseCase) GenerateChildCode(ctx context.Context, userID string, createdIndex int) (string, error) {
	code := helper.GenerateChildCode(createdIndex)

	ownerCode := &model.OwnerCode{
		OwnerID:      userID,
		OwnerRole:    string(constants.ChildRole),
		Code:         code,
		CreatedIndex: createdIndex,
	}

	err := s.ownerCodeRepo.Create(ctx, ownerCode)
	if err != nil {
		return "", err
	}

	_ = s.cachingService.SetChildCode(ctx, userID, code)
	return code, nil
}
