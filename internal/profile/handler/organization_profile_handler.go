package handler

import (
	"errors"
	"net/http"
	"profile-service/helper"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/service"

	"github.com/gin-gonic/gin"
)

type OrganizationProfileHandler struct {
	service service.OrganizationProfileService
}

func NewOrganizationProfileHandler(service service.OrganizationProfileService) *OrganizationProfileHandler {
	return &OrganizationProfileHandler{service: service}
}

func (h *OrganizationProfileHandler) UploadSummary(c *gin.Context) {
	var req request.UploadOrganizationSummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	organizationID := c.Param("organization_id")
	if organizationID == "" {
		helper.SendError(c, http.StatusBadRequest, errors.New("organization_id is required"), helper.ErrInvalidRequest)
		return
	}
	req.OrganizationID = organizationID

	err := h.service.UploadSummary(c.Request.Context(), req.OrganizationID, req.Summary)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "Success", nil)
}
