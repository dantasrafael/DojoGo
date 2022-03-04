package model

import (
	"financial/src/domain/entity"
)

type RegisterAccountNotification struct {
	ID     uint64                  `json:"id"`
	Status entity.EnrollmentStatus `json:"status"`
}
