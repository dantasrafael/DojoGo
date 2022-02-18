package entities

import (
	"fmt"
	"modulo-escolar/src/core/utils/messages"
	"time"
)

type Enrollment struct {
	ID           uint64           `json:"id"`
	Student      Student          `json:"student"`
	Course       Course           `json:"course"`
	Installments uint8            `json:"installments"`
	Status       EnrollmentStatus `json:"status"`
	CreatedAt    time.Time        `json:"createdAt"`
}

func (enrollment *Enrollment) Prepare() error {
	if err := enrollment.validate(); err != nil {
		return err
	}

	return nil
}

func (enrollment *Enrollment) validate() error {
	if enrollment.Student.ID == 0 {
		return fmt.Errorf(messages.REQUIRED_FIELD, "Estudante")
	}

	if enrollment.Course.ID == 0 {
		return fmt.Errorf(messages.REQUIRED_FIELD, "Curso")
	}

	if enrollment.Installments == 0 {
		return fmt.Errorf(messages.REQUIRED_FIELD, "Parcelas")
	}

	return nil
}
