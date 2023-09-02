package model

type CommonStudentsRequest struct {
	TeacherEmails []string `form:"teacher" binding:"required,dive,email"`
}