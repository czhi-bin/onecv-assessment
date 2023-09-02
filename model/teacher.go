package model

// Teacher is the model for teacher table
type Teacher struct {
	ID    int64  `gorm:"primaryKey;autoIncrement"`
	Email string `gorm:"unique;not null;type:varchar(50)"`
}
