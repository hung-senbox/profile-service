package route

import (
	"profile-service/internal/gateway"
	"profile-service/internal/middleware"
	"profile-service/internal/profile/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(r *gin.Engine, ho *handler.OwnerCodeHandler, oh *handler.OrganizationProfileHandler, userGw gateway.UserGateway) {
	apiV1 := r.Group("/api/v1")

	adminGroup := apiV1.Group("/admin")
	adminGroup.Use(middleware.Secured(userGw))
	{
		profileAdmin := adminGroup.Group("/profiles")
		{
			// organization profile
			organizationProfile := profileAdmin.Group("/organization")
			{
				organizationProfile.POST("/summary", oh.UploadSummary)
				organizationProfile.GET("/summary", oh.GetSummary)
			}
		}
	}

	ownerCodePublic := apiV1.Group("/gateway/profiles/owner-code")
	{
		ownerCodePublic.POST("/user/generate", ho.GenerateUserCode)
	}

	gatewayGroup := apiV1.Group("/gateway")
	gatewayGroup.Use(middleware.Secured(userGw))
	{
		profilesAdmin := gatewayGroup.Group("/profiles")
		{
			ownerCode := profilesAdmin.Group("/owner-code")
			{
				ownerCode.POST("/student/generate", ho.GenerateStudentCode)
				ownerCode.POST("/teacher/generate", ho.GenerateTeacherCode)
				ownerCode.POST("/staff/generate", ho.GenerateStaffCode)
				ownerCode.POST("/parent/generate", ho.GenerateParentCode)
				ownerCode.POST("/child/generate", ho.GenerateChildCode)
				ownerCode.POST("/device/generate", ho.GenerateDeviceCode)
				ownerCode.POST("/organization/generate", ho.GenerateOrganizationCode)

				ownerCode.GET("/student/:student_id", ho.GetStudentCode)
				ownerCode.GET("/teacher/:teacher_id", ho.GetTeacherCode)
				ownerCode.GET("/staff/:staff_id", ho.GetStaffCode)
				ownerCode.GET("/parent/:parent_id", ho.GetParentCode)
				ownerCode.GET("/user/:user_id", ho.GetUserCode)
				ownerCode.GET("/child/:child_id", ho.GetChildCode)
				ownerCode.GET("/device/:device_id", ho.GetDeviceCode)
				ownerCode.GET("/organization/:organization_id", ho.GetOrganizationCode)
			}
		}
	}
}
