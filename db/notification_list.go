package db

import (
	"errors"

	"github.com/czhi-bin/onecv-assessment/model"
	"github.com/czhi-bin/onecv-assessment/utils"
)

// Retrieve the non-suspended registered students of the teacher
func GetNonSuspendedRegisteredStudents(teacherEmail string) ([]string, error) {
	teacherId, err := getTeacherId(teacherEmail)
	if err != nil {
		if err.Error() == "teacher not found" {
			return nil, nil
		}
		return nil, err
	}

	var students []model.Student
	err = DB.Model(&model.Registration{}).
		Select("students.id, students.email").Where("teacher_id = ?", teacherId).
		Joins("JOIN students ON registrations.student_id = students.id AND students.is_suspended = false").
		Find(&students).Error
	if err != nil {
		utils.Logger.Error(err, teacherId, "Error in retrieving non-suspended registered student from database")
		return nil, err
	}

	if len(students) == 0 {
		// no non-suspended students registered to the teacher
		return nil, nil
	}

	studentEmails:= getStudentEmails(students)

	return studentEmails, nil
}

// Retrieve the teacher ID by email from DB
func getTeacherId(teacherEmail string) (int64, error) {
	var teacher model.Teacher
	db := DB.Where("email = ?", teacherEmail).Limit(1).Find(&teacher)
	if db.Error != nil {
		utils.Logger.Error(db.Error, teacherEmail, "Error in retrieving teacher using email from database")
		return 0, db.Error
	} else if db.RowsAffected == 0 {
		return 0, errors.New("teacher not found")
	}

	return teacher.ID, nil
}

func getStudentEmails(students []model.Student) []string {
	studentEmails := make([]string, len(students))
	for i, student := range students {
		studentEmails[i] = student.Email
	}

	return studentEmails
}

func CheckIsSuspended(studentEmail string) (bool, error) {
	db := DB.Where("email = ? AND is_suspended = ?", studentEmail, true).Find(&model.Student{})
	if db.Error != nil {
		utils.Logger.Error(db.Error, studentEmail, 
			"Error in checking student suspension status using email from database")
		return false, db.Error
	}

	// if the student is suspended, db.RowsAffected will be 1, will return True
	return db.RowsAffected == 1, nil
}
