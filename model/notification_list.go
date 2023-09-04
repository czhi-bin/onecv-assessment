package model

type NotificationListRequest struct {
	TeacherEmail string `json:"teacher"  binding:"required,email"` // Teacher email
	Notification string `json:"notification" binding:"required"`   // The notification message and the students tagged
}
