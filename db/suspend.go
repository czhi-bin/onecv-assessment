package db

import (
	"errors"

	"github.com/czhi-bin/onecv-assessment/model"
	"github.com/czhi-bin/onecv-assessment/utils"
)

func UpdateStudentSuspendStatus(email string, status bool) error {
	db := DB.Model(&model.Student{}).Where("email = ?", email).Update("is_suspended", status)
	if db.RowsAffected == 0 {
		return errors.New("student not found")
	}

	if db.Error != nil {
		utils.Logger.Error(db.Error, email, status, "Error in updating student suspend status in database")
		return db.Error
	}
	
	return nil
}