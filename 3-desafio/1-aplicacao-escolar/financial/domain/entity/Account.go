package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	ID        uuid.UUID
	CreatedAt time.Time
	ClientID  string
	CourseID  string
	Total     float32
}
