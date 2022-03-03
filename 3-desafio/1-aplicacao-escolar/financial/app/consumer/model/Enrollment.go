package model

import uuid "github.com/satori/go.uuid"

type Enrollment struct {
	ID           uuid.UUID
	StudentID    string  `json:"student_id"`
	CourseID     string  `json:"course_id"`
	Installments uint8   `json:"installments"`
	Total        float32 `json:"total"`
}
