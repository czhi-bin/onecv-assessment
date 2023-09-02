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

// // Create a registration record in DB
// func createRegistration(registration *model.Registration) error {
// 	err := DB.Create(registration).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Retrieve the student ID in DB by email, if not found, create one
// func getStudentIDByEmail(email string) (int64, error) {
// 	var student model.Student
// 	err := DB.Where("email = ?", email).Limit(1).Find(&student).Error
// 	if err != nil {
// 		// error in finding student
// 		return 0, err
// 	}

// 	// student not found, create one record for the student
// 	if student == (model.Student{}) {
// 		student = model.Student{Email: email}
// 		err := createStudent(&student)
// 		if err != nil {
// 			// error in creating student
// 			return 0, err
// 		}
// 	}

// 	// either student is found or created, return the student ID
// 	return student.ID, nil
// }

// // Create a student record in DB and return the created student ID
// func createStudent(student *model.Student) error {
// 	err := DB.Create(student).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Retrieve the teacher ID in DB by email, if not found, create one
// func getTeacherIDByEmail(email string) (int64, error) {
// 	var teacher model.Teacher
// 	err := DB.Where("email = ?", email).Limit(1).Find(&teacher).Error
// 	if err != nil {
// 		// error in finding teacher
// 		return 0, err
// 	}

// 	// teacher not found, create one record for the teacher
// 	if teacher == (model.Teacher{}) {
// 		teacher = model.Teacher{Email: email}
// 		err := createTeacher(&teacher)
// 		if err != nil {
// 			// error in creating teacher
// 			return 0, err
// 		}
// 	}

// 	// either teacher is found or created, return the teacher ID
// 	return teacher.ID, nil
// }

// // Create a teacher record in DB and return the created teacher ID
// func createTeacher(teacher *model.Teacher) error {
// 	err := DB.Create(teacher).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }