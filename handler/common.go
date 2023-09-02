package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/model"
)

// @router /api/commonstudents [GET]
func GetCommonStudentList(c *gin.Context) {
	// TODO: implement this
	// idea: loop through each teacher, find the students registered
	// to each of them, then do a set intersection to find the common

	var err error
	var req model.CommonStudentsRequest

	// Check the query parameters
	err = c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request! Please provide the correct query parameters.",
		})
		return
	}

	// Get the common students
	commonStudents, err := db.GetCommonStudents(req.TeacherEmails)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve the common students! Please try again.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"students": commonStudents,
	})
}

