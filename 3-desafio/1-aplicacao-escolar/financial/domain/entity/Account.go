package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	ID           uuid.UUID
	ClientID     string
	CourseID     string
	ExternalID   string
	Installments uint8
	Total        float64
	Status       EnrollmentStatus
	CreatedAt    time.Time
}
