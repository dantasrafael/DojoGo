package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/producers"
	"modulo-escolar/src/infra/repositories"
)

func DeleteEnrollmentUsecase(id *uint64) error {
	if err := repositories.DeleteEnrollment(id); err != nil {
		return err
	}

	producers.DeleteEnrollmentProducer(&entities.Enrollment{ID: *id})
	return nil
}
