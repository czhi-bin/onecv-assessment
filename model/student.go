package model

// Student is the model for student table
type Student struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Email       string `gorm:"unique;not null;type:varchar(50)"`
	IsSuspended bool   `gorm:"default:false;not null"`
}
