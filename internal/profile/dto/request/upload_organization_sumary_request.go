package request

type UploadOrganizationSummaryRequest struct {
	OrganizationID string `json:"organization_id"`
	Summary        string `json:"summary" binding:"required"`
}
