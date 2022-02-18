package entities

import (
	"fmt"
	"modulo-escolar/src/core/utils/messages"
	"strings"
	"time"
)

type Student struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"createdAt"`
}

func (student *Student) Prepare() error {
	if err := student.validate(); err != nil {
		return err
	}

	student.format()
	return nil
}

func (student *Student) validate() error {
	if student.Name == "" {
		return fmt.Errorf(messages.REQUIRED_FIELD, "Nome")
	}

	if student.Email == "" {
		return fmt.Errorf(messages.REQUIRED_FIELD, "E-mail")
	}

	return nil
}

func (student *Student) format() {
	student.Name = strings.TrimSpace(student.Name)
	student.Email = strings.TrimSpace(student.Email)
}
