package service

import (
	"context"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/usecase"
)

type OrganizationProfileService interface {
	UploadSummary(ctx context.Context, req request.UploadOrganizationSummaryRequest) error
	GetSummary(ctx context.Context) (string, error)
}

type organizationProfileService struct {
	organizationProfileUsecase usecase.OrganizationProfileUsecase
}

func NewOrganizationProfileService(organizationProfileUsecase usecase.OrganizationProfileUsecase) OrganizationProfileService {
	return &organizationProfileService{organizationProfileUsecase: organizationProfileUsecase}
}

func (s *organizationProfileService) UploadSummary(ctx context.Context, req request.UploadOrganizationSummaryRequest) error {
	return s.organizationProfileUsecase.UploadSummary(ctx, req)
}

func (s *organizationProfileService) GetSummary(ctx context.Context) (string, error) {
	return s.organizationProfileUsecase.GetSummary(ctx)
}
