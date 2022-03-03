package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func UpdateEnrollmentStatusUsecase(message interface{}) error {
	model := message.(*entities.Enrollment)
	if err := repositories.UpdateEnrollmentStatus(&model.ID, &model.Status); err != nil {
		return err
	}

	return nil
}
