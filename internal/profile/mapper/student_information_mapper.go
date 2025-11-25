package mapper

import (
	"profile-service/internal/profile/dto/response"
	"profile-service/internal/profile/model"
)

func ToStudentInformationResponse4Web(studentInfo *model.StudentInformation) *response.StudentInformationResponse {
	return &response.StudentInformationResponse{
		StudentID:         studentInfo.StudentID,
		DOB:               studentInfo.DOB,
		Gender:            studentInfo.Gender,
		StudyLevel:        studentInfo.StudyLevel,
		MinWaterMustDrink: studentInfo.MinWaterMustDrink,
		Description:       studentInfo.Description,
		Mode:              studentInfo.Mode,
	}
}

func ToStudentInformationResponse4Gw(studentInfo *model.StudentInformation) *response.StudentInformationResponse {
	return &response.StudentInformationResponse{
		StudentID:         studentInfo.StudentID,
		DOB:               studentInfo.DOB,
		Gender:            studentInfo.Gender,
		StudyLevel:        studentInfo.StudyLevel,
		MinWaterMustDrink: studentInfo.MinWaterMustDrink,
		Description:       studentInfo.Description,
		Mode:              studentInfo.Mode,
	}
}
