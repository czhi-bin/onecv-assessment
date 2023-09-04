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

func TestRegisterStudent_ValidRequest(t *testing.T) {
    // Set up the test database
    db.Init(true)
	
    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/register", RegisterStudent)

    // Create a fake request
    payload := `{"teacher": "teacherken@gmail.com", 
				"students": 
					[
			        "studentjon@gmail.com", 
		            "studenthon@gmail.com"
					]
				}`
    req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Should return 204 No Content, with an empty body
    assert.Equal(t, http.StatusNoContent, w.Code)
    assert.Empty(t, w.Body.String())

    // // Check if the rows are inserted into database properly
    var teacher model.Teacher
    var student1 model.Student
    var student2 model.Student
    assert.NoError(t, db.DB.Where("email = ?", "teacherken@gmail.com").First(&teacher).Error)
    assert.NoError(t, db.DB.Where("email = ?", "studentjon@gmail.com").First(&student1).Error)
    assert.NoError(t, db.DB.Where("email = ?", "studenthon@gmail.com").First(&student2).Error)
    assert.NoError(t, db.DB.Where("teacher_id = ? AND student_id = ?", teacher.ID, student1.ID).
                    First(&model.Registration{}).Error)
    assert.NoError(t, db.DB.Where("teacher_id = ? AND student_id = ?", teacher.ID, student2.ID).
                    First(&model.Registration{}).Error)
}

func TestRegisterStudent_MissingTeacher(t *testing.T) {
    // Set up the test database
    db.Init(true)
	
    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/register", RegisterStudent)

    // Create a fake request
    payload := `{"students": 
					[
			        "studentjon@gmail.com", 
		            "studenthon@gmail.com"
					]
				}`
    req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Should return 400 Bad Request, with an error message
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Equal(t, `{"message":"Invalid request body! Please check your request body and try again."}`, w.Body.String())
}

func TestRegisterStudent_MissingStudent(t *testing.T) {
    // Set up the test database
    db.Init(true)
	
    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/register", RegisterStudent)

    // Create a fake request
    payload := `{"teacher": "teacherken@gmail.com"}`
    req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Should return 400 Bad Request, with an error message
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Equal(t, `{"message":"Invalid request body! Please check your request body and try again."}`, w.Body.String())
}

func TestRegisterStudent_InvalidEmail(t *testing.T) {
    // Set up the test database
    db.Init(true)
	
    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/register", RegisterStudent)

    // Create a fake request
    payload := `{"teacher": "teacherken", 
				"students": 
					[
			        "studentjon", 
		            "studenthon@gmail.com"
					]
				}`
    req, _ := http.NewRequest("POST", "/api/register", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Should return 400 Bad Request, with an error message
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Equal(t, `{"message":"Invalid request body! Please check your request body and try again."}`, w.Body.String())
}