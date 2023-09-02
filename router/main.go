package router

import (
	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/onecv-assessment/handler"
)

// Register all the routes
func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	apiGroup.POST("/register", handler.Register)

	apiGroup.GET("/commonstudents", handler.GetCommonStudentList)

	apiGroup.POST("/suspend", handler.Suspend)

	apiGroup.GET("/retrievefornotifications", handler.GetNotificationList)
}
