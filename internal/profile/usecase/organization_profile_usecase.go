package usecase

import (
	"context"
	"errors"
	gw_response "profile-service/internal/gateway/dto/response"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/model"
	"profile-service/internal/profile/repository"
	"profile-service/pkg/constants"
)

type OrganizationProfileUsecase interface {
	UploadSummary(ctx context.Context, req request.UploadOrganizationSummaryRequest) error
	GetSummary(ctx context.Context) (string, error)
}

type organizationProfileUsecase struct {
	organizationProfileRepository repository.OrganizationProfileRepository
}

func NewOrganizationProfileUsecase(organizationProfileRepository repository.OrganizationProfileRepository) OrganizationProfileUsecase {
	return &organizationProfileUsecase{organizationProfileRepository}
}

func (s *organizationProfileUsecase) UploadSummary(ctx context.Context, req request.UploadOrganizationSummaryRequest) error {

	currentUser, ok := ctx.Value(constants.CurrentUserKey).(*gw_response.CurrentUser)
	if !ok {
		return errors.New("current user not found")
	}

	if currentUser.IsSuperAdmin {
		return errors.New("you are not allowed to upload summary")
	}

	if currentUser.OrganizationAdmin == nil {
		return errors.New("organization admin not found")
	}

	organizationProfile := &model.OrganizationProfile{
		OrganizationID: currentUser.OrganizationAdmin.ID,
		Summary:        req.Summary,
	}
	return s.organizationProfileRepository.UploadSummary(ctx, organizationProfile)
}

func (s *organizationProfileUsecase) GetSummary(ctx context.Context) (string, error) {
	currentUser, ok := ctx.Value(constants.CurrentUserKey).(*gw_response.CurrentUser)
	if !ok {
		return "", errors.New("current user not found")
	}

	return s.organizationProfileRepository.GetSummary(ctx, currentUser.OrganizationAdmin.ID)
}
