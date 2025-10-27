package service

import (
	"context"
	"profile-service/internal/gateway"
	"profile-service/internal/profile/repository"
	"profile-service/internal/profile/usecase"
)

type OwnerCodeService interface {
	// generate owner codes
	GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error)
	GenerateTeacherCode(ctx context.Context, teacherID string, createdIndex int) (string, error)
	GenerateStaffCode(ctx context.Context, staffID string, createdIndex int) (string, error)
	GenerateParentCode(ctx context.Context, parentID string, createdIndex int) (string, error)
	GenerateUserCode(ctx context.Context, userID string, createdIndex int) (string, error)
	GenerateChildCode(ctx context.Context, childID string, createdIndex int) (string, error)

	// get owner codes
	GetStudentCodeByStudentID(ctx context.Context, studentID string) (*string, error)
	GetTeacherCodeByTeacherID(ctx context.Context, teacherID string) (*string, error)
	GetStaffCodeByStaffID(ctx context.Context, staffID string) (*string, error)
	GetParentCodeByParentID(ctx context.Context, parentID string) (*string, error)
	GetUserCodeByUserID(ctx context.Context, userID string) (*string, error)
	GetChildCodeByChildID(ctx context.Context, childID string) (*string, error)
}

type ownerCodeService struct {
	ownerCodeRepo            repository.OwnerCodeRepository
	cachedUserGw             gateway.UserGateway
	generateOwnerCodeUseCase usecase.GenerateOwnerCodeUseCase
	getOwnerCodeUseCase      usecase.GetOwnerCodeUseCase
}

func NewOwnerCodeService(
	ownerCodeRepo repository.OwnerCodeRepository,
	cachedUserGw gateway.UserGateway,
	generateOwnerCodeUseCase usecase.GenerateOwnerCodeUseCase,
	getOwnerCodeUseCase usecase.GetOwnerCodeUseCase,
) OwnerCodeService {
	return &ownerCodeService{
		ownerCodeRepo:            ownerCodeRepo,
		cachedUserGw:             cachedUserGw,
		generateOwnerCodeUseCase: generateOwnerCodeUseCase,
		getOwnerCodeUseCase:      getOwnerCodeUseCase,
	}
}

func (s *ownerCodeService) GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error) {
	return s.generateOwnerCodeUseCase.GenerateStudentCode(ctx, studentID, createdIndex)
}

func (s *ownerCodeService) GenerateTeacherCode(ctx context.Context, teacherID string, createdIndex int) (string, error) {
	return s.generateOwnerCodeUseCase.GenerateTeacherCode(ctx, teacherID, createdIndex)
}

func (s *ownerCodeService) GenerateStaffCode(ctx context.Context, staffID string, createdIndex int) (string, error) {
	return s.generateOwnerCodeUseCase.GenerateStaffCode(ctx, staffID, createdIndex)
}

func (s *ownerCodeService) GenerateParentCode(ctx context.Context, parentID string, createdIndex int) (string, error) {
	return s.generateOwnerCodeUseCase.GenerateParentCode(ctx, parentID, createdIndex)
}

func (s *ownerCodeService) GenerateUserCode(ctx context.Context, userID string, createdIndex int) (string, error) {
	return s.generateOwnerCodeUseCase.GenerateUserCode(ctx, userID, createdIndex)
}

func (s *ownerCodeService) GenerateChildCode(ctx context.Context, childID string, createdIndex int) (string, error) {
	return s.generateOwnerCodeUseCase.GenerateChildCode(ctx, childID, createdIndex)
}

func (s *ownerCodeService) GetStudentCodeByStudentID(ctx context.Context, studentID string) (*string, error) {
	return s.getOwnerCodeUseCase.GetStudentCodeByStudentID(ctx, studentID)
}

func (s *ownerCodeService) GetTeacherCodeByTeacherID(ctx context.Context, teacherID string) (*string, error) {
	return s.getOwnerCodeUseCase.GetTeacherCodeByTeacherID(ctx, teacherID)
}

func (s *ownerCodeService) GetStaffCodeByStaffID(ctx context.Context, staffID string) (*string, error) {
	return s.getOwnerCodeUseCase.GetStaffCodeByStaffID(ctx, staffID)
}

func (s *ownerCodeService) GetParentCodeByParentID(ctx context.Context, parentID string) (*string, error) {
	return s.getOwnerCodeUseCase.GetParentCodeByParentID(ctx, parentID)
}

func (s *ownerCodeService) GetUserCodeByUserID(ctx context.Context, userID string) (*string, error) {
	return s.getOwnerCodeUseCase.GetUserCodeByUserID(ctx, userID)
}

func (s *ownerCodeService) GetChildCodeByChildID(ctx context.Context, childID string) (*string, error) {
	return s.getOwnerCodeUseCase.GetChildCodeByChildID(ctx, childID)
}
