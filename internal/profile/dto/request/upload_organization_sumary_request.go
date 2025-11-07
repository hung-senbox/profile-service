package request

type UploadOrganizationSummaryRequest struct {
	Summary string `json:"summary" binding:"required"`
}
