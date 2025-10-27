package service

import (
	"context"
	"profile-service/internal/gateway"
	"profile-service/internal/profile/repository"
	"profile-service/internal/profile/usecase"
)

type OwnerCodeService interface {
	GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error)
	GenerateTeacherCode(ctx context.Context, teacherID string, createdIndex int) (string, error)
	GenerateStaffCode(ctx context.Context, staffID string, createdIndex int) (string, error)
	GenerateParentCode(ctx context.Context, parentID string, createdIndex int) (string, error)
	GenerateUserCode(ctx context.Context, userID string, createdIndex int) (string, error)
	GenerateChildCode(ctx context.Context, childID string, createdIndex int) (string, error)
}

type ownerCodeService struct {
	ownerCodeRepo            repository.OwnerCodeRepository
	cachedUserGw             gateway.UserGateway
	generateOwnerCodeUseCase usecase.GenerateOwnerCodeUseCase
}

func NewOwnerCodeService(ownerCodeRepo repository.OwnerCodeRepository, cachedUserGw gateway.UserGateway, generateOwnerCodeUseCase usecase.GenerateOwnerCodeUseCase) OwnerCodeService {
	return &ownerCodeService{
		ownerCodeRepo:            ownerCodeRepo,
		cachedUserGw:             cachedUserGw,
		generateOwnerCodeUseCase: generateOwnerCodeUseCase,
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
