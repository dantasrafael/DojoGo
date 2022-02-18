package entities

import (
	"fmt"
	"modulo-escolar/src/core/utils/messages"
	"strings"
	"time"
)

type Course struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
}

func (course *Course) Prepare() error {
	if err := course.validate(); err != nil {
		return err
	}

	course.format()
	return nil
}

func (course *Course) validate() error {
	if course.Name == "" {
		return fmt.Errorf(messages.REQUIRED_FIELD, "Nome")
	}

	if course.Value == 0.0 {
		return fmt.Errorf(messages.REQUIRED_FIELD, "Valor")
	}

	return nil
}

func (course *Course) format() {
	course.Name = strings.TrimSpace(course.Name)
}
