package model

// RegisterRequest is the request body for registering students to a teacher.
type RegisterRequest struct {
	TeacherEmail  string   `json:"teacher"  binding:"required,email"`      // Teacher email
	StudentEmails []string `json:"students" binding:"required,dive,email"` // Student emails
}

// Registration is the many-to-many relationship between Teacher and Student.
// gorm creates a table named "registrations" with composite primary key (teacher_id, student_id)
// which references the primary keys of "teachers" and "students" tables.
type Registration struct {
	TeacherID int64 `gorm:"primaryKey;autoIncrement:false;"`
	StudentID int64 `gorm:"primaryKey;autoIncrement:false;"`
	Teacher   Teacher
	Student   Student
}
