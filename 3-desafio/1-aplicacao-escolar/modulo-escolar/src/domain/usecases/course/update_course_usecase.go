package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func UpdateCourseUsecase(id *uint64, model *entities.Course) error {
	if err := repositories.UpdateCourse(id, model); err != nil {
		return err
	}

	return nil
}
