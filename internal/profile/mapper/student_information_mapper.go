package mapper

import (
	"profile-service/internal/profile/dto/response"
	"profile-service/internal/profile/model"
)

func ToStudentInformationResponse4Web(studentInfo *model.StudentInformation) *response.StudentInformationResponse {
	gender := "boy"
	if studentInfo.Gender {
		gender = "girl"
	}
	return &response.StudentInformationResponse{
		DOB:               studentInfo.DOB,
		Gender:            gender,
		StudyLevel:        studentInfo.StudyLevel,
		MinWaterMustDrink: studentInfo.MinWaterMustDrink,
		Description:       studentInfo.Description,
		Mode:              studentInfo.Mode,
	}
}

func ToStudentInformationResponse4Gw(studentInfo *model.StudentInformation) *response.StudentInformationResponse {
	gender := "boy"
	if studentInfo.Gender {
		gender = "girl"
	}
	return &response.StudentInformationResponse{
		DOB:               studentInfo.DOB,
		Gender:            gender,
		StudyLevel:        studentInfo.StudyLevel,
		MinWaterMustDrink: studentInfo.MinWaterMustDrink,
		Description:       studentInfo.Description,
		Mode:              studentInfo.Mode,
	}
}
