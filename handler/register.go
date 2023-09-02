package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/model"
)

// @router /api/register [POST]
func Register(c *gin.Context) {
	var err error
	var req model.RegisterRequest
	
	// Check request body
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body! Please check your request body and try again.",
		})
		return
	}

	// Register the student(s) to the teacher
	for _, studentEmail := range req.StudentEmails {
		err = db.CreateRegistration(req.TeacherEmail, studentEmail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to register the student(s) to the teacher! Please try again.",
			})
			return
		}
	}

	// Registration successful
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Successfully registered the students to the teacher!",
	})
}

