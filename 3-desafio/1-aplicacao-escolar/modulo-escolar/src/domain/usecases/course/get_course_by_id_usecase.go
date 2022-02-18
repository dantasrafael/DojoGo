package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func GetCourseByIdUsecase(id *uint64) (*entities.Course, error) {
	record, err := repositories.FindCourseById(id)
	if err != nil {
		return nil, err
	}

	return record, nil
}
