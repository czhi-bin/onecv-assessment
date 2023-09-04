package db

import (
	"errors"

	"github.com/czhi-bin/onecv-assessment/model"
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
		return nil, err
	}

	if len(students) == 0 {
		// no non-suspended students registered to the teacher
		return nil, nil
	}

	studentEmails, err := getStudentEmails(students)
	if err != nil {
		return nil, err
	}

	return studentEmails, nil
}

// Retrieve the teacher ID by email from DB
func getTeacherId(teacherEmail string) (int64, error) {
	var teacher model.Teacher
	db := DB.Where("email = ?", teacherEmail).Limit(1).Find(&teacher)
	if db.Error != nil {
		return 0, db.Error
	} else if db.RowsAffected == 0 {
		return 0, errors.New("teacher not found")
	}

	return teacher.ID, nil
}

func getStudentEmails(students []model.Student) ([]string, error) {
	studentEmails := make([]string, len(students))
	for i, student := range students {
		studentEmails[i] = student.Email
	}

	return studentEmails, nil
}

func CheckIsSuspended(studentEmail string) (bool, error) {
	db := DB.Where("email = ? AND is_suspended = ?", studentEmail, true).Find(&model.Student{})
	if db.Error != nil {
		return false, db.Error
	}

	// if the student is suspended, db.RowsAffected will be 1, will return True
	return db.RowsAffected == 1, nil
}
