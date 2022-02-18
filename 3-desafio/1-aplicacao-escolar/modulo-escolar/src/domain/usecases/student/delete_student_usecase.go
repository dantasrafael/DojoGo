package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/producers"
	"modulo-escolar/src/infra/repositories"
)

func DeleteStudentUsecase(id *uint64) error {
	if err := repositories.DeleteStudent(id); err != nil {
		return err
	}

	producers.DeleteStudentProducer(&entities.Student{ID: *id})
	return nil
}
