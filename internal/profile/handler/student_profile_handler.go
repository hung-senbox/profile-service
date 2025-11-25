package handler

import (
	"net/http"
	"profile-service/helper"
	"profile-service/internal/profile/dto/request"
	"profile-service/internal/profile/service"

	"github.com/gin-gonic/gin"
)

type StudentProfileHandler struct {
	service service.StudentProfileService
}

func NewStudentProfileHandler(service service.StudentProfileService) *StudentProfileHandler {
	return &StudentProfileHandler{service: service}
}

func (h *StudentProfileHandler) UploadStudentInfo(c *gin.Context) {
	var req request.UploadStudentInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}
	err := h.service.UploadStudentInfo(c.Request.Context(), req)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}
	helper.SendSuccess(c, http.StatusOK, "Success", nil)
}

func (h *StudentProfileHandler) GetStudentInfo4Web(c *gin.Context) {
	studentID := c.Param("student_id")
	if studentID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, helper.ErrInvalidRequest)
		return
	}
	studentInfo, err := h.service.GetStudentInfo4Web(c.Request.Context(), studentID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}
	helper.SendSuccess(c, http.StatusOK, "Success", studentInfo)
}

func (h *StudentProfileHandler) GetStudentInfo4Gw(c *gin.Context) {
	studentID := c.Param("student_id")
	if studentID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, helper.ErrInvalidRequest)
		return
	}
	studentInfo, err := h.service.GetStudentInfo4Gw(c.Request.Context(), studentID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}
	helper.SendSuccess(c, http.StatusOK, "Success", studentInfo)
}
