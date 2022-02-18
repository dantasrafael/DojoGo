package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func GetAllCoursesUsecase(name *string) ([]entities.Course, error) {
	list, err := repositories.FindAllCourses(name)
	if err != nil {
		return nil, err
	}

	return list, nil
}
