package service

import (
	"context"
	"profile-service/internal/profile/usecase"
)

type OrganizationProfileService interface {
	UploadSummary(ctx context.Context, organizationID, summary string) error
}

type organizationProfileService struct {
	organizationProfileUsecase usecase.OrganizationProfileUsecase
}

func NewOrganizationProfileService(organizationProfileUsecase usecase.OrganizationProfileUsecase) OrganizationProfileService {
	return &organizationProfileService{organizationProfileUsecase: organizationProfileUsecase}
}

func (s *organizationProfileService) UploadSummary(ctx context.Context, organizationID, summary string) error {
	return s.organizationProfileUsecase.UploadSummary(ctx, organizationID, summary)
}
