package service

import (
	"context"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/dto/response"
	"profile-service/internal/profile/usecase"
)

type StudentProfileService interface {
	UploadStudentInfo(ctx context.Context, req request.UploadStudentInfoRequest) error
	GetStudentInfo4Web(ctx context.Context, studentID string) (*response.StudentProfileResponse, error)
	GetStudentInfo4Gw(ctx context.Context, studentID string) (*response.StudentProfileResponse, error)
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

func (s *studentProfileService) GetStudentInfo4Web(ctx context.Context, studentID string) (*response.StudentProfileResponse, error) {
	studentInfo, err := s.studentInformationUsecase.GetStudentInfo4Web(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return &response.StudentProfileResponse{
		StudentInformation: *studentInfo,
	}, nil
}

func (s *studentProfileService) GetStudentInfo4Gw(ctx context.Context, studentID string) (*response.StudentProfileResponse, error) {
	studentInfo, err := s.studentInformationUsecase.GetStudentInfo4Gw(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return &response.StudentProfileResponse{
		StudentInformation: *studentInfo,
	}, nil
}
