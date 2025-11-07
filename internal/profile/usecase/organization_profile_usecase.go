package usecase

import (
	"context"
	"profile-service/internal/profile/model"
	"profile-service/internal/profile/repository"
)

type OrganizationProfileUsecase interface {
	UploadSummary(ctx context.Context, organizationID, summary string) error
}

type organizationProfileUsecase struct {
	organizationProfileRepository repository.OrganizationProfileRepository
}

func NewOrganizationProfileUsecase(organizationProfileRepository repository.OrganizationProfileRepository) OrganizationProfileUsecase {
	return &organizationProfileUsecase{organizationProfileRepository}
}

func (s *organizationProfileUsecase) UploadSummary(ctx context.Context, organizationID, summary string) error {
	organizationProfile := &model.OrganizationProfile{
		OrganizationID: organizationID,
		Summary:        summary,
	}
	return s.organizationProfileRepository.UploadSummary(ctx, organizationProfile)
}
