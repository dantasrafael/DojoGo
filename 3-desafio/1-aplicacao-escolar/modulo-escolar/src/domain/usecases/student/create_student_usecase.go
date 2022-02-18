package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func CreateStudentUsecase(model *entities.Student) error {
	if err := repositories.CreateStudent(model); err != nil {
		return err
	}

	return nil
}
