package db

import (
	"github.com/czhi-bin/onecv-assessment/model"
	"github.com/czhi-bin/onecv-assessment/utils"
)

type CommonStudent struct {
	ID    int64
	Email string
}

// Retrieve the common students of the teachers
func GetCommonStudents(teacherEmails []string) ([]string, error) {
	var teachers []model.Teacher

	// retrieve teacher IDs by emails
	err := DB.Where("email IN (?)", teacherEmails).Find(&teachers).Error
	if err != nil {
		utils.Logger.Error(err, teacherEmails, "Error in retrieving teachers using teacher IDs from database")
		return nil, err
	} else if len(teachers) != len(teacherEmails) {
		// some teacher have no students, common students will be empty
		return nil, nil
	}

	// retrieve students registered to the teachers
	var commonStudents []CommonStudent
	teacherIds := getTeacherIds(teachers)
	err = DB.Model(&model.Registration{}).
		Select("student_id as id").Where("teacher_id IN (?)", teacherIds).
		Group("student_id").Having("count(teacher_id) = ?", len(teachers)).
		Find(&commonStudents).Error
	if err != nil {
		utils.Logger.Error(err, teacherIds, len(teachers), "Error in retrieving registrations from database")
		return nil, err
	}

	commonStudentEmails, err := getStudentEmailsFromIds(commonStudents)
	if err != nil {
		return nil, err
	}

	return commonStudentEmails, nil
}

// Map the teacher to its teacher ID
func getTeacherIds(teachers []model.Teacher) []int64 {
	teacherIds := make([]int64, len(teachers))
	for i, teacher := range teachers {
		teacherIds[i] = teacher.ID
	}

	return teacherIds
}

func getStudentEmailsFromIds(commonStudents []CommonStudent) ([]string, error) {
	studentEmails := make([]string, len(commonStudents))
	for i, commonStudent := range commonStudents {
		err := DB.Model(&model.Student{}).Where("id = ?", commonStudent.ID).Find(&commonStudent).Error
		if err != nil {
			utils.Logger.Error(err, commonStudent, "Error in retrieving student email using ID from database")
			return nil, err
		}
		studentEmails[i] = commonStudent.Email
	}

	return studentEmails, nil
}
