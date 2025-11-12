package usecase

import (
	"context"
	"profile-service/helper"
	"profile-service/internal/profile/model"
	"profile-service/internal/profile/repository"
	"profile-service/pkg/constants"

	profile_caching_service "github.com/hung-senbox/senbox-cache-service/pkg/cache/caching"
)

type GenerateOwnerCodeUseCase interface {
	GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error)
	GenerateTeacherCode(ctx context.Context, teacherID string, createdIndex int) (string, error)
	GenerateStaffCode(ctx context.Context, staffID string, createdIndex int) (string, error)
	GenerateParentCode(ctx context.Context, parentID string, createdIndex int) (string, error)
	GenerateUserCode(ctx context.Context, userID string, createdIndex int) (string, error)
	GenerateChildCode(ctx context.Context, userID string, createdIndex int) (string, error)
	GenerateDeviceCode(ctx context.Context, deviceID string, createdIndex int) (string, error)
}

type generateOwnerCodeUseCase struct {
	ownerCodeRepo         repository.OwnerCodeRepository
	profileCachingService profile_caching_service.CachingProfileService
}

func NewGenerateOwnerCodeUseCase(ownerCodeRepo repository.OwnerCodeRepository, profileCachingService profile_caching_service.CachingProfileService) GenerateOwnerCodeUseCase {
	return &generateOwnerCodeUseCase{
		ownerCodeRepo:         ownerCodeRepo,
		profileCachingService: profileCachingService,
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

	_ = s.profileCachingService.SetStudentCode(ctx, studentID, code)
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

	_ = s.profileCachingService.SetTeacherCode(ctx, teacherID, code)
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

	_ = s.profileCachingService.SetStaffCode(ctx, staffID, code)
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

	_ = s.profileCachingService.SetParentCode(ctx, parentID, code)
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

	_ = s.profileCachingService.SetUserCode(ctx, userID, code)
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

	_ = s.profileCachingService.SetChildCode(ctx, userID, code)
	return code, nil
}

func (s *generateOwnerCodeUseCase) GenerateDeviceCode(ctx context.Context, deviceID string, createdIndex int) (string, error) {
	code := helper.GenerateDeviceCode(createdIndex)

	ownerCode := &model.OwnerCode{
		OwnerID:      deviceID,
		OwnerRole:    string(constants.DeviceRole),
		Code:         code,
		CreatedIndex: createdIndex,
	}

	err := s.ownerCodeRepo.Create(ctx, ownerCode)
	if err != nil {
		return "", err
	}

	_ = s.profileCachingService.SetDeviceCode(ctx, deviceID, code)
	return code, nil
}
