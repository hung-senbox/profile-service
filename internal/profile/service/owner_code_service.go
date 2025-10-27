package service

import (
	"context"
	"profile-service/internal/gateway"
	"profile-service/internal/profile/model"
	"profile-service/internal/profile/repository"
	"profile-service/pkg/constants"
	"strconv"
)

type OwnerCodeService interface {
	GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error)
}

type ownerCodeService struct {
	ownerCodeRepo repository.OwnerCodeRepository
	cachedUserGw  gateway.UserGateway
}

func NewOwnerCodeService(ownerCodeRepo repository.OwnerCodeRepository, cachedUserGw gateway.UserGateway) OwnerCodeService {
	return &ownerCodeService{
		ownerCodeRepo: ownerCodeRepo,
		cachedUserGw:  cachedUserGw,
	}
}

func (s *ownerCodeService) GenerateStudentCode(ctx context.Context, studentID string, createdIndex int) (string, error) {
	code := studentID + "_" + strconv.Itoa(createdIndex)

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

	return code, nil
}
