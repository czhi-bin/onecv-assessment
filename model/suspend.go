package model

// SuspendRequest is the request body for suspending a student.
type SuspendRequest struct {
	StudentEmail string `json:"student" binding:"required,email"`
}