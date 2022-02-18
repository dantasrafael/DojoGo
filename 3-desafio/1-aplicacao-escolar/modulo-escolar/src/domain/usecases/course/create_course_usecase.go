package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func CreateCourseUsecase(model *entities.Course) error {
	if err := repositories.CreateCourse(model); err != nil {
		return err
	}

	return nil
}
