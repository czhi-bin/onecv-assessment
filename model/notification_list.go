package model

type NotificationListRequest struct {
	TeacherEmail  string   `json:"teacher"  binding:"required,email"`      		// Teacher email
	Notification 	string `json:"notification" binding:"required,dive,email"`  // The notification message and the students tagged
}