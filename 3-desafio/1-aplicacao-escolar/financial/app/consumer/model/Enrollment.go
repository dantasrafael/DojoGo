package model

type Enrollment struct {
	StudentID    string  `json:"student_id"`
	CourseID     string  `json:"course_id"`
	Installments uint8   `json:"installments"`
	Total        float32 `json:"total"`
}
