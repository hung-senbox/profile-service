package usecase

import (
	"context"
	"errors"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/dto/response"
	"profile-service/internal/profile/mapper"
	"profile-service/internal/profile/model"
	"profile-service/internal/profile/repository"
	"profile-service/pkg/constants"
	"time"
)

type StudentInformationUsecase interface {
	UploadStudentInfo(ctx context.Context, req request.UploadStudentInfoRequest) error
	GetStudentInfo4Web(ctx context.Context, studentID string) (*response.StudentInformationResponse, error)
	GetStudentInfo4Gw(ctx context.Context, studentID string) (*response.StudentInformationResponse, error)
}

type studentInformationUsecase struct {
	studentInformationRepository repository.StudentInformationRepository
}

func NewStudentInformationUsecase(studentInformationRepository repository.StudentInformationRepository) StudentInformationUsecase {
	return &studentInformationUsecase{studentInformationRepository}
}

func (s *studentInformationUsecase) UploadStudentInfo(ctx context.Context, req request.UploadStudentInfoRequest) error {

	if err := s.validateStudentInfo(req); err != nil {
		return err
	}

	studentInformation := &model.StudentInformation{
		StudentID:         req.StudentID,
		DOB:               req.DOB,
		Gender:            req.Gender,
		StudyLevel:        req.StudyLevel,
		MinWaterMustDrink: req.MinWaterMustDrink,
		Description:       req.Description,
		Mode:              req.Mode,
	}
	return s.studentInformationRepository.Create(ctx, studentInformation)
}

func (s *studentInformationUsecase) GetStudentInfo(ctx context.Context, studentID string) (*model.StudentInformation, error) {
	return s.studentInformationRepository.GetByStudentID(ctx, studentID)
}

func (s *studentInformationUsecase) validateStudentInfo(req request.UploadStudentInfoRequest) error {
	// validate request
	// dob
	if req.DOB.After(time.Now()) {
		return errors.New("dob is in the future")
	}
	// gender
	if req.Gender != string(constants.GenderMale) && req.Gender != string(constants.GenderFemale) {
		return errors.New("gender is invalid")
	}
	// mode
	if req.Mode != string(constants.TeacherMode) && req.Mode != string(constants.ParentMode) && req.Mode != string(constants.StudentMode) && req.Mode != string(constants.ClassroomMode) && req.Mode != string(constants.OrganizationMode) {
		return errors.New("mode is invalid")
	}
	// study level 1 - 10
	if req.StudyLevel < 1 || req.StudyLevel > 10 {
		return errors.New("study level is invalid")
	}
	return nil
}

func (s *studentInformationUsecase) GetStudentInfo4Gw(ctx context.Context, studentID string) (*response.StudentInformationResponse, error) {
	studentInfo, err := s.GetStudentInfo(ctx, studentID)
	if err != nil {
		return nil, err
	}

	return mapper.ToStudentInformationResponse4Gw(studentInfo), nil
}

func (s *studentInformationUsecase) GetStudentInfo4Web(ctx context.Context, studentID string) (*response.StudentInformationResponse, error) {
	studentInfo, err := s.GetStudentInfo(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return mapper.ToStudentInformationResponse4Web(studentInfo), nil
}
