package db

import (
	"github.com/czhi-bin/onecv-assessment/model"
)

// Register the student to the teacher
func CreateRegistration(teacherEmail, studentEmail string) error {
	// retrieve teacher ID by email
	var teacher model.Teacher
	err := DB.Where(model.Teacher{Email: teacherEmail}).FirstOrCreate(&teacher).Error
	if err != nil {
		// error in retrieving teacher
		return err
	}

	// retrieve student ID by email
	var student model.Student
	err = DB.Where(model.Student{Email: studentEmail}).FirstOrCreate(&student).Error
	if err != nil {
		// error in retrieving student
		return err
	}

	// create registration record if not exists
	var registration model.Registration
	err = DB.Where(model.Registration{
		TeacherID: teacher.ID, 
		StudentID: student.ID,
	}).FirstOrCreate(&registration).Error
	if err != nil {
		// error in creating registration
		return err
	}

	// registration created successfully
	return nil
}