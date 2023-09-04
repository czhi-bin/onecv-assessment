package db

import (
	"github.com/czhi-bin/onecv-assessment/model"
	"github.com/czhi-bin/onecv-assessment/utils"
)

// Register the student to the teacher
func CreateRegistration(teacherEmail, studentEmail string) error {
	// retrieve teacher ID by email
	var teacher model.Teacher
	err := DB.Where(model.Teacher{Email: teacherEmail}).FirstOrCreate(&teacher).Error
	if err != nil {
		utils.Logger.Error(err, teacherEmail, "Error in retrieving/creating teacher record using email from database")
		return err
	}

	// retrieve student ID by email
	var student model.Student
	err = DB.Where(model.Student{Email: studentEmail}).FirstOrCreate(&student).Error
	if err != nil {
		utils.Logger.Error(err, studentEmail, "Error in retrieving/creating student record using email from database")
		return err
	}

	// create registration record if not exists
	var registration model.Registration
	err = DB.Where(model.Registration{
		TeacherID: teacher.ID, 
		StudentID: student.ID,
	}).FirstOrCreate(&registration).Error
	if err != nil {
		utils.Logger.Error(err, teacher.ID, student.ID, "Error in creating registration record in database")
		return err
	}

	// registration created successfully
	return nil
}