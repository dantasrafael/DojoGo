package model

import (
	"financial/domain/entity"
	"time"
)

type Enrollment struct {
	ID           uint64                  `json:"id"`
	Student      Student                 `json:"student"`
	Course       Course                  `json:"course"`
	Installments uint8                   `json:"installments"`
	Status       entity.EnrollmentStatus `json:"status"`
	CreatedAt    time.Time               `json:"createdAt"`
}

type Student struct {
	ID uint64 `json:"id"`
}

type Course struct {
	ID    uint64  `json:"id"`
	Value float64 `json:"value"`
}
