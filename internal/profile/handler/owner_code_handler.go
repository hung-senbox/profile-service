package handler

import (
	"net/http"
	"profile-service/helper"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/service"

	"github.com/gin-gonic/gin"
)

type OwnerCodeHandler struct {
	service service.OwnerCodeService
}

func NewOwnerCodeHandler(service service.OwnerCodeService) *OwnerCodeHandler {
	return &OwnerCodeHandler{
		service: service,
	}
}

func (h *OwnerCodeHandler) GenerateStudentCode(c *gin.Context) {
	var req request.GenerateOwnerCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	res, err := h.service.GenerateStudentCode(c.Request.Context(), req.OwnerID, req.CreatedIndex)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Success", res)
}
