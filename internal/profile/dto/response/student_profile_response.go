package response

import "time"

type StudentProfileResponse struct {
	StudentInformation StudentInformationResponse `json:"student_information"`
}

type StudentInformationResponse struct {
	DOB               time.Time `json:"dob"`
	Gender            bool      `json:"gender"`
	StudyLevel        uint      `json:"study_level"`
	MinWaterMustDrink uint      `json:"min_water_must_drink"`
	Description       string    `json:"description"`
	Mode              string    `json:"mode"`
}
