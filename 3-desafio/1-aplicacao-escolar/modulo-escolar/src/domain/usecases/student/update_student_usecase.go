package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func UpdateStudentUsecase(id *uint64, model *entities.Student) error {
	if err := repositories.UpdateStudent(id, model); err != nil {
		return err
	}

	return nil
}
