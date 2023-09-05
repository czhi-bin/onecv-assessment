package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/czhi-bin/onecv-assessment/db"
	"github.com/czhi-bin/onecv-assessment/model"
)

func TestGetCommonStudentList_ValidRequest1(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForCommonStudentList()

    // Initialize a new Gin router
    router := gin.Default()

    router.GET("/api/commonstudents", GetCommonStudentList)

    // Create a fake request
    req, _ := http.NewRequest("GET", "/api/commonstudents", nil)
    params := url.Values{}
    params.Add("teacher", "teacherwow@gmail.com")
    req.URL.RawQuery = params.Encode()

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with a list of common students, and the student only under teacher
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, 
        `{"students":["commonstudent1@gmail.com","commonstudent2@gmail.com","student_only_under_teacher_wow@gmail.com"]}`, 
        w.Body.String())
}

func TestGetCommonStudentList_ValidRequest2(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForCommonStudentList()

    // Initialize a new Gin router
    router := gin.Default()

    router.GET("/api/commonstudents", GetCommonStudentList)

    // Create a fake request
    req, _ := http.NewRequest("GET", "/api/commonstudents", nil)
    params := url.Values{}
    params.Add("teacher", "teacherwow@gmail.com")
    params.Add("teacher", "teacherwowza@gmail.com")
    req.URL.RawQuery = params.Encode()

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with a list of common students
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, 
        `{"students":["commonstudent1@gmail.com","commonstudent2@gmail.com"]}`, 
        w.Body.String())
}

func TestGetCommonStudentList_TeacherWithoutStudent(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForCommonStudentList()

    // Initialize a new Gin router
    router := gin.Default()

    router.GET("/api/commonstudents", GetCommonStudentList)

    // Create a fake request
    req, _ := http.NewRequest("GET", "/api/commonstudents", nil)
    params := url.Values{}
    params.Add("teacher", "teacherwow@gmail.com")
    params.Add("teacher", "teachernostudent@gmail.com")
    req.URL.RawQuery = params.Encode()

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 200 OK, with an empty list since there will be no common students
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, `{"students":[]}`, w.Body.String())
}

func TestGetCommonStudentList_InvalidEmail(t *testing.T) {
	// Set up the test database
    db.Init(true)
	populateTestDatabaseForCommonStudentList()

    // Initialize a new Gin router
    router := gin.Default()

    router.GET("/api/commonstudents", GetCommonStudentList)

    // Create a fake request
    req, _ := http.NewRequest("GET", "/api/commonstudents", nil)
    params := url.Values{}
    params.Add("teacher", "teacherwow")
    params.Add("teacher", "teachernostudent@gmail.com")
    req.URL.RawQuery = params.Encode()

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    // Should return 400 Bad Request, with an error message
    assert.Equal(t, http.StatusBadRequest, w.Code)
    assert.Equal(t, `{"message":"Invalid request! Please provide the correct query parameters."}`, w.Body.String())
}

func populateTestDatabaseForCommonStudentList() {
    // Populate the test database
    var teacher1 model.Teacher
    var teacher2 model.Teacher
    var commonStudent1 model.Student
    var commonStudent2 model.Student
    var wowStudent model.Student
    _ = db.DB.Where(&model.Teacher{Email: "teacherwow@gmail.com"}).FirstOrCreate(&teacher1).Error
    _ = db.DB.Where(&model.Teacher{Email: "teacherwowza@gmail.com"}).FirstOrCreate(&teacher2).Error
    _ = db.DB.Where(&model.Student{Email: "commonstudent1@gmail.com"}).FirstOrCreate(&commonStudent1).Error
    _ = db.DB.Where(&model.Student{Email: "commonstudent2@gmail.com"}).FirstOrCreate(&commonStudent2).Error
    _ = db.DB.Where(&model.Student{Email: "student_only_under_teacher_wow@gmail.com"}).FirstOrCreate(&wowStudent).Error
    _ = db.DB.Where(&model.Registration{TeacherID: teacher1.ID, StudentID: commonStudent1.ID}).FirstOrCreate(&model.Registration{}).Error
    _ = db.DB.Where(&model.Registration{TeacherID: teacher1.ID, StudentID: commonStudent2.ID}).FirstOrCreate(&model.Registration{}).Error
    _ = db.DB.Where(&model.Registration{TeacherID: teacher1.ID, StudentID: wowStudent.ID}).FirstOrCreate(&model.Registration{}).Error
    _ = db.DB.Where(&model.Registration{TeacherID: teacher2.ID, StudentID: commonStudent1.ID}).FirstOrCreate(&model.Registration{}).Error
    _ = db.DB.Where(&model.Registration{TeacherID: teacher2.ID, StudentID: commonStudent2.ID}).FirstOrCreate(&model.Registration{}).Error
}