package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/model"
)

// @router /api/suspend [POST]
func SuspendStudent(c *gin.Context) {
	var err error
	var req model.SuspendRequest

	// Check request body
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body! Please check your request body and try again.",
		})
		return
	}

	// Suspend the student
	err = db.UpdateStudentSuspendStatus(req.StudentEmail, true)
	if err != nil {
		// Student does not exist
		if err.Error() == "student not found" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Student not found! Please check the student email and try again.",
			})
			return
		}

		// Other errors
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to suspend the student! Please try again.",
		})
		return
	}

	// Suspend successful
	c.JSON(http.StatusNoContent, nil)
}
