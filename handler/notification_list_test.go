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

func TestGetNotificationList_ValidRequest1(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "teachernotification@gmail.com", 
				"notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with a list of registered students and mentioned students
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, 
        `{"recipients":["studentagnes@gmail.com","studentbob@gmail.com","studentmiche@gmail.com"]}`, 
        w.Body.String())
}

func TestGetNotificationList_ValidRequest2(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "teachernotification@gmail.com", 
				"notification": "Hey everybody!"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with a list of registered students
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"recipients":["studentbob@gmail.com"]}`, w.Body.String())
}

func TestGetNotificationList_AllStudentSuspended(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "allstudentsuspended@gmail.com", 
				"notification": "Hey everybody!"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with an empty list
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"recipients":[]}`, w.Body.String())
}

func TestGetNotificationList_NoRegisteredStudent1(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "nostudent@gmail.com", 
				"notification": "Hey everybody!"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with an empty list
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"recipients":[]}`, w.Body.String())
}

func TestGetNotificationList_NoRegisteredStudent2(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "nostudent@gmail.com", 
				"notification": "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with a list of mentioned students
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"recipients":["studentagnes@gmail.com","studentmiche@gmail.com"]}`, w.Body.String())
}

func TestGetNotificationList_MentionedStudentSuspended(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "allstudentsuspended@gmail.com", 
				"notification": "Hello students! @studentsuspended@gmail.com"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with an empty list
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"recipients":[]}`, w.Body.String())
}

func TestGetNotificationList_SomeMentionedStudentSuspended(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "allstudentsuspended@gmail.com", 
				"notification": "Hello students! @studentsuspended@gmail.com @studenthon@gmail.com"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with only one of the mentioned students
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"recipients":["studenthon@gmail.com"]}`, w.Body.String())
}

func TestGetNotificationList_InvalidEmail(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForNotificationList()

    // Initialize a new Gin router
    router := gin.Default()

    router.POST("/api/retrievefornotifications", GetNotificationList)

    // Create a fake request
	payload := `{"teacher": "teacherken", 
				"notification": "Hey everybody!"
				}`
    req, _ := http.NewRequest("POST", "/api/retrievefornotifications", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 400 Bad Request, with an error message
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Equal(t, `{"message":"Invalid request body! Please check your request body and try again."}`, w.Body.String())
}

func populateTestDatabaseForNotificationList() {
    // Populate the test database
    var teacher model.Teacher
	var allStudentSuspended model.Teacher
    var registeredStudent model.Student
	var suspendedStudent model.Student
    _ = db.DB.Where(&model.Teacher{Email: "teachernotification@gmail.com"}).FirstOrCreate(&teacher).Error
	_ = db.DB.Where(&model.Teacher{Email: "allstudentsuspended@gmail.com"}).FirstOrCreate(&allStudentSuspended).Error
    _ = db.DB.Where(&model.Student{Email: "studentbob@gmail.com"}).FirstOrCreate(&registeredStudent).Error
	_ = db.DB.Where(&model.Student{Email: "studentsuspended@gmail.com", IsSuspended: true}).FirstOrCreate(&suspendedStudent).Error
    _ = db.DB.Where(&model.Student{Email: "studentagnes@gmail.com"}).FirstOrCreate(&model.Student{}).Error
    _ = db.DB.Where(&model.Student{Email: "studentmiche@gmail.com"}).FirstOrCreate(&model.Student{}).Error
    _ = db.DB.Where(&model.Registration{TeacherID: teacher.ID, StudentID: registeredStudent.ID}).FirstOrCreate(&model.Registration{}).Error
	_ = db.DB.Where(&model.Registration{TeacherID: allStudentSuspended.ID, StudentID: suspendedStudent.ID}).FirstOrCreate(&model.Registration{}).Error
}