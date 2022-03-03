package usecases

import (
	"encoding/json"
	"log"
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func UpdateEnrollmentStatusUsecase(message string) error {
	log.Printf("processing message: %s", message)
	var model entities.Enrollment
	if err := json.Unmarshal([]byte(message), &model); err != nil {
		return err
	}

	if err := repositories.UpdateEnrollmentStatus(&model.ID, &model.Status); err != nil {
		return err
	}

	return nil
}
