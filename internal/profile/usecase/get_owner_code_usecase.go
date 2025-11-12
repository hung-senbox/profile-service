package usecase

import (
	"context"
	"profile-service/internal/profile/repository"
	"profile-service/pkg/constants"
)

type GetOwnerCodeUseCase interface {
	GetOwnerCodeByOwnerIDAndRole(ctx context.Context, ownerID string, ownerRole constants.OwnerRoles) (*string, error)
	GetStudentCodeByStudentID(ctx context.Context, studentID string) (*string, error)
	GetTeacherCodeByTeacherID(ctx context.Context, teacherID string) (*string, error)
	GetStaffCodeByStaffID(ctx context.Context, staffID string) (*string, error)
	GetParentCodeByParentID(ctx context.Context, parentID string) (*string, error)
	GetUserCodeByUserID(ctx context.Context, userID string) (*string, error)
	GetChildCodeByChildID(ctx context.Context, childID string) (*string, error)
	GetDeviceCodeByDeviceID(ctx context.Context, deviceID string) (*string, error)
}

type getOwnerCodeUseCase struct {
	ownerCodeRepo repository.OwnerCodeRepository
}

func NewGetOwnerCodeUseCase(ownerCodeRepo repository.OwnerCodeRepository) GetOwnerCodeUseCase {
	return &getOwnerCodeUseCase{
		ownerCodeRepo: ownerCodeRepo,
	}
}

func (s *getOwnerCodeUseCase) GetOwnerCodeByOwnerIDAndRole(ctx context.Context, ownerID string, ownerRole constants.OwnerRoles) (*string, error) {
	ownerCode, err := s.ownerCodeRepo.FindByOwnerIDAndRole(ctx, ownerID, string(ownerRole))
	if err != nil {
		return nil, err
	}
	return &ownerCode.Code, nil
}

func (s *getOwnerCodeUseCase) GetStudentCodeByStudentID(ctx context.Context, studentID string) (*string, error) {
	return s.GetOwnerCodeByOwnerIDAndRole(ctx, studentID, constants.StudentRole)
}

func (s *getOwnerCodeUseCase) GetTeacherCodeByTeacherID(ctx context.Context, teacherID string) (*string, error) {
	return s.GetOwnerCodeByOwnerIDAndRole(ctx, teacherID, constants.TeacherRole)
}

func (s *getOwnerCodeUseCase) GetStaffCodeByStaffID(ctx context.Context, staffID string) (*string, error) {
	return s.GetOwnerCodeByOwnerIDAndRole(ctx, staffID, constants.StaffRole)
}

func (s *getOwnerCodeUseCase) GetParentCodeByParentID(ctx context.Context, parentID string) (*string, error) {
	return s.GetOwnerCodeByOwnerIDAndRole(ctx, parentID, constants.ParentRole)
}

func (s *getOwnerCodeUseCase) GetUserCodeByUserID(ctx context.Context, userID string) (*string, error) {
	return s.GetOwnerCodeByOwnerIDAndRole(ctx, userID, constants.UserRole)
}

func (s *getOwnerCodeUseCase) GetChildCodeByChildID(ctx context.Context, childID string) (*string, error) {
	return s.GetOwnerCodeByOwnerIDAndRole(ctx, childID, constants.ChildRole)
}

func (s *getOwnerCodeUseCase) GetDeviceCodeByDeviceID(ctx context.Context, deviceID string) (*string, error) {
	return s.GetOwnerCodeByOwnerIDAndRole(ctx, deviceID, constants.DeviceRole)
}
