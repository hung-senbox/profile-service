package request

type UploadStudentInfoRequest struct {
	StudentID         string `json:"student_id"`
	DOB               string `json:"dob"`
	Gender            uint   `json:"gender"`
	StudyLevel        uint   `json:"study_level"`
	MinWaterMustDrink uint   `json:"min_water_must_drink"`
	Description       string `json:"description"`
	Mode              string `json:"mode"`
}
