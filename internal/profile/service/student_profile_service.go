package service

import (
	"context"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/dto/response"
	"profile-service/internal/profile/usecase"
)

type StudentProfileService interface {
	UploadStudentInfo(ctx context.Context, req request.UploadStudentInfoRequest) error
	GetStudentProfile4Web(ctx context.Context, studentID string) (*response.StudentProfileResponse, error)
	GetStudentProfile4Gw(ctx context.Context, studentID string) (*response.StudentProfileResponse, error)
}

type studentProfileService struct {
	studentInformationUsecase usecase.StudentInformationUsecase
}

func NewStudentProfileService(studentInformationUsecase usecase.StudentInformationUsecase) StudentProfileService {
	return &studentProfileService{studentInformationUsecase}
}

func (s *studentProfileService) UploadStudentInfo(ctx context.Context, req request.UploadStudentInfoRequest) error {
	return s.studentInformationUsecase.UploadStudentInfo(ctx, req)
}

func (s *studentProfileService) GetStudentProfile4Web(ctx context.Context, studentID string) (*response.StudentProfileResponse, error) {
	studentInfo, err := s.studentInformationUsecase.GetStudentProfile4Web(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return &response.StudentProfileResponse{
		StudentInformation: *studentInfo,
	}, nil
}

func (s *studentProfileService) GetStudentProfile4Gw(ctx context.Context, studentID string) (*response.StudentProfileResponse, error) {
	studentInfo, err := s.studentInformationUsecase.GetStudentProfile4Gw(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return &response.StudentProfileResponse{
		StudentInformation: *studentInfo,
	}, nil
}
