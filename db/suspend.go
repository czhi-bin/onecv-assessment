package db

import (
	"errors"

	"github.com/czhi-bin/onecv-assessment/model"
)

func UpdateStudentSuspendStatus(email string, status bool) error {
	db := DB.Model(&model.Student{}).Where("email = ?", email).Update("is_suspended", status)
	if db.RowsAffected == 0 {
		// return ErrStudentNotFound
		return errors.New("student not found")
	}

	err := db.Error
	return err
}