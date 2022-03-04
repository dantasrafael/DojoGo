package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Invoice struct {
	ID          uuid.UUID
	Account     *Account
	Installment uint8
	CreatedAt   time.Time
	PaidAt      time.Time
	DueDate     time.Time
	Value       float64
}
