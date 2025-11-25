package response

import "time"

type StudentInformationResponse struct {
	StudentID         string    `json:"student_id"`
	DOB               time.Time `json:"dob"`
	Gender            string    `json:"gender"`
	StudyLevel        uint      `json:"study_level"`
	MinWaterMustDrink uint      `json:"min_water_must_drink"`
	Description       string    `json:"description"`
	Mode              string    `json:"mode"`
}
