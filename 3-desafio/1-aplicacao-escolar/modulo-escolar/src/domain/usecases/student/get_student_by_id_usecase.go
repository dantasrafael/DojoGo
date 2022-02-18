package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func GetStudentByIdUsecase(id *uint64) (*entities.Student, error) {
	record, err := repositories.FindStudentById(id)
	if err != nil {
		return nil, err
	}

	return record, nil
}
