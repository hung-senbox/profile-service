package route

import (
	"profile-service/internal/gateway"
	"profile-service/internal/middleware"
	"profile-service/internal/profile/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(r *gin.Engine, ho *handler.OwnerCodeHandler, userGw gateway.UserGateway) {
	// Admin routes
	gatewayGroup := r.Group("/api/v1/gateway")
	gatewayGroup.Use(middleware.Secured(userGw))
	{
		profilesAdmin := gatewayGroup.Group("/profiles")
		{
			// owner code routes
			ownerCode := profilesAdmin.Group("/owner-code")
			{
				ownerCode.POST("/student/generate", ho.GenerateStudentCode)
				ownerCode.POST("/teacher/generate", ho.GenerateTeacherCode)
				ownerCode.POST("/staff/generate", ho.GenerateStaffCode)
				ownerCode.POST("/parent/generate", ho.GenerateParentCode)
				ownerCode.POST("/user/generate", ho.GenerateUserCode)
				ownerCode.POST("/child/generate", ho.GenerateChildCode)

				ownerCode.GET("/student/:student_id", ho.GetStudentCode)
				ownerCode.GET("/teacher/:teacher_id", ho.GetTeacherCode)
				ownerCode.GET("/staff/:staff_id", ho.GetStaffCode)
				ownerCode.GET("/parent/:parent_id", ho.GetParentCode)
				ownerCode.GET("/user/:user_id", ho.GetUserCode)
				ownerCode.GET("/child/:child_id", ho.GetChildCode)
			}
		}
	}
}
