package router

import (
	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/onecv-assessment/handler"
	"github.com/czhi-bin/onecv-assessment/utils"
)

// Register all the routes
func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")

	apiGroup.POST("/register", handler.RegisterStudent)

	apiGroup.GET("/commonstudents", handler.GetCommonStudentList)

	apiGroup.POST("/suspend", handler.SuspendStudent)

	apiGroup.POST("/retrievefornotifications", handler.GetNotificationList)

	utils.Logger.Info("Routes registered")
}