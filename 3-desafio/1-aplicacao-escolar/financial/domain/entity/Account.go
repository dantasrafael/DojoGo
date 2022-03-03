package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	ID           uuid.UUID
	ClientID     string
	CourseID     string
	Installments uint8
	Total        float32
	CreatedAt    time.Time
}
