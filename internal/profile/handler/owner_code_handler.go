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

func (h *OwnerCodeHandler) GenerateTeacherCode(c *gin.Context) {
	var req request.GenerateOwnerCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	res, err := h.service.GenerateTeacherCode(c.Request.Context(), req.OwnerID, req.CreatedIndex)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Success", res)
}

func (h *OwnerCodeHandler) GenerateStaffCode(c *gin.Context) {
	var req request.GenerateOwnerCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	res, err := h.service.GenerateStaffCode(c.Request.Context(), req.OwnerID, req.CreatedIndex)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Success", res)
}

func (h *OwnerCodeHandler) GenerateParentCode(c *gin.Context) {
	var req request.GenerateOwnerCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	res, err := h.service.GenerateParentCode(c.Request.Context(), req.OwnerID, req.CreatedIndex)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Success", res)
}

func (h *OwnerCodeHandler) GenerateUserCode(c *gin.Context) {
	var req request.GenerateOwnerCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	res, err := h.service.GenerateUserCode(c.Request.Context(), req.OwnerID, req.CreatedIndex)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Success", res)
}

func (h *OwnerCodeHandler) GenerateChildCode(c *gin.Context) {
	var req request.GenerateOwnerCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.SendError(c, http.StatusBadRequest, err, helper.ErrInvalidRequest)
		return
	}

	res, err := h.service.GenerateChildCode(c.Request.Context(), req.OwnerID, req.CreatedIndex)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Success", res)
}

// ==================== Get Owner Codes =====================
func (h *OwnerCodeHandler) GetStudentCode(c *gin.Context) {
	studentID := c.Param("student_id")
	if studentID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, "student_id is required")
		return
	}

	res, err := h.service.GetStudentCodeByStudentID(c.Request.Context(), studentID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "Success", res)
}

func (h *OwnerCodeHandler) GetTeacherCode(c *gin.Context) {
	teacherID := c.Param("teacher_id")
	if teacherID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, "teacher_id is required")
		return
	}

	res, err := h.service.GetTeacherCodeByTeacherID(c.Request.Context(), teacherID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "Success", res)
}

func (h *OwnerCodeHandler) GetStaffCode(c *gin.Context) {
	staffID := c.Param("staff_id")
	if staffID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, "staff_id is required")
		return
	}

	res, err := h.service.GetStaffCodeByStaffID(c.Request.Context(), staffID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "Success", res)
}

func (h *OwnerCodeHandler) GetParentCode(c *gin.Context) {
	parentID := c.Param("parent_id")
	if parentID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, "parent_id is required")
		return
	}
	res, err := h.service.GetParentCodeByParentID(c.Request.Context(), parentID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "Success", res)
}

func (h *OwnerCodeHandler) GetUserCode(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, "user_id is required")
		return
	}

	res, err := h.service.GetUserCodeByUserID(c.Request.Context(), userID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "Success", res)
}

func (h *OwnerCodeHandler) GetChildCode(c *gin.Context) {
	childID := c.Param("child_id")
	if childID == "" {
		helper.SendError(c, http.StatusBadRequest, nil, "child_id is required")
		return
	}

	res, err := h.service.GetChildCodeByChildID(c.Request.Context(), childID)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, err, helper.ErrInvalidOperation)
		return
	}

	helper.SendSuccess(c, http.StatusOK, "Success", res)
}
