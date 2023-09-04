package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/model"
)

func TestSuspendStudent_ValidRequest(t *testing.T) {
	// Set up the test database
    db.Init(true)
	// Add the student record to the database
    _ = db.DB.Where(&model.Student{Email: "studentmary@gmail.com"}).FirstOrCreate(&model.Student{}).Error

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/suspend", SuspendStudent)

    // Create a fake request
    payload := `{
					"student": "studentmary@gmail.com"
				}`
    req, _ := http.NewRequest("POST", "/api/suspend", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Should return 204 No Content, with an empty body
    assert.Equal(t, http.StatusNoContent, w.Code)
    assert.Empty(t, w.Body.String())

	// // Check if the row are updated properly
    assert.NoError(t, db.DB.Where("email = ? AND is_suspended = ?", "studentmary@gmail.com", true).First(&model.Student{}).Error)
}

func TestSuspendStudent_InvalidEmail(t *testing.T) {
	// Set up the test database
    db.Init(true)

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/suspend", SuspendStudent)

    // Create a fake request
    payload := `{
					"student": "studentmary"
				}`
    req, _ := http.NewRequest("POST", "/api/suspend", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Should return 400 Bad Request, with an error message
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Equal(t, `{"message":"Invalid request body! Please check your request body and try again."}`, w.Body.String())
}

func TestSuspendStudent_NonExistentStudent(t *testing.T) {
	// Set up the test database
    db.Init(true)

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/suspend", SuspendStudent)

    // Create a fake request
    payload := `{
					"student": "studentnoone@gmail.com"
				}`
    req, _ := http.NewRequest("POST", "/api/suspend", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Should return 400 Bad Request, with an error message
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Equal(t, `{"message":"Student not found! Please check the student email and try again."}`, w.Body.String())
}