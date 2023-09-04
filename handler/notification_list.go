package handler

import (
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/model"
)

// @router /api/retrievefornotifications [GET]
func GetNotificationList(c *gin.Context) {
	var err error
	var req model.NotificationListRequest

	// Check request body
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body! Please check your request body and try again.",
		})
		return
	}

	// Extract out the mentioned student from the notification message
	studentEmails := make(map[string]struct{})
	for _, studentEmail := range strings.Fields(req.Notification) {
		if studentEmail[0] == '@' {
			isSuspended, err := db.CheckIsSuspended(studentEmail[1:])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Error in checking suspension status! Please try again later.",
				})
				return
			} else if isSuspended {
				continue
			}
			studentEmails[studentEmail[1:]] = struct{}{}
		}
	}

	// Get the list of students who are registered under the teacher
	registeredStudents, err := db.GetNonSuspendedRegisteredStudents(req.TeacherEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve the registered students! Please try again later.",
		})
		return
	}

	// Using a map to eliminate duplicates
	for _, studentEmail := range registeredStudents {
		studentEmails[studentEmail] = struct{}{}
	}

	if len(studentEmails) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"recipients": []string{},
		})
		return
	}

	// Convert the map to a slice
	studentEmailsSlice := make([]string, len(studentEmails))
	i := 0
	for studentEmail := range studentEmails {
		studentEmailsSlice[i] = studentEmail
		i++
	}

	c.JSON(http.StatusOK, gin.H{
		"recipients": studentEmailsSlice,
	})
}
