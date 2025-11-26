package mapper

import (
	"profile-service/internal/profile/dto/response"
	"profile-service/internal/profile/model"
)

func ToStudentInformationResponse4Web(studentInfo *model.StudentInformation) *response.StudentInformationResponse {
	var dobStr *string
	if studentInfo.DOB != nil {
		formatted := studentInfo.DOB.Format("2006-01-02")
		dobStr = &formatted
	}
	return &response.StudentInformationResponse{
		DOB:               dobStr,
		Gender:            studentInfo.Gender,
		StudyLevel:        studentInfo.StudyLevel,
		MinWaterMustDrink: studentInfo.MinWaterMustDrink,
		Description:       studentInfo.Description,
		Mode:              studentInfo.Mode,
	}
}

func ToStudentInformationResponse4Gw(studentInfo *model.StudentInformation) *response.StudentInformationResponse {
	var dobStr *string
	if studentInfo.DOB != nil {
		formatted := studentInfo.DOB.Format("2006-01-02")
		dobStr = &formatted
	}
	return &response.StudentInformationResponse{
		DOB:               dobStr,
		Gender:            studentInfo.Gender,
		StudyLevel:        studentInfo.StudyLevel,
		MinWaterMustDrink: studentInfo.MinWaterMustDrink,
		Description:       studentInfo.Description,
		Mode:              studentInfo.Mode,
	}
}
