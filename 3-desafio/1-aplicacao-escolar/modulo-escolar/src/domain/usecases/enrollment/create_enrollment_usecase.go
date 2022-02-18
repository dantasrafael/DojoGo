package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/producers"
	"modulo-escolar/src/infra/repositories"
)

func CreateEnrollmentUsecase(model *entities.Enrollment) error {
	model.Status = entities.ADIMPLENTE
	if err := repositories.CreateEnrollment(model); err != nil {
		return err
	}

	producers.CreateEnrollmentProducer(model)
	return nil
}
