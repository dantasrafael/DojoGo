package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/producers"
	"modulo-escolar/src/infra/repositories"
)

func DeleteCourseUsecase(id *uint64) error {
	if err := repositories.DeleteCourse(id); err != nil {
		return err
	}

	producers.DeleteCourseProducer(&entities.Course{ID: *id})
	return nil
}
