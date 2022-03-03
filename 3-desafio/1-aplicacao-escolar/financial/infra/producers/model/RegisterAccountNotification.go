package model

import "financial/domain/entity"

type RegisterAccountNotification struct {
	ID     uint64                  `json:"id"`
	Status entity.EnrollmentStatus `json:"status"`
}
