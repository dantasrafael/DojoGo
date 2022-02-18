package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func GetAllStudentsUsecase(name *string) ([]entities.Student, error) {
	list, err := repositories.FindAllStudents(name)
	if err != nil {
		return nil, err
	}

	return list, nil
}
