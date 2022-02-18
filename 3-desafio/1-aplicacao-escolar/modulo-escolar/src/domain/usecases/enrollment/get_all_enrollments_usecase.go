package usecases

import (
	"modulo-escolar/src/domain/entities"
	"modulo-escolar/src/infra/repositories"
)

func GetAllEnrollmentsUsecase(studentName, courseName *string) ([]entities.Enrollment, error) {
	list, err := repositories.FindAllEnrollments(studentName, courseName)
	if err != nil {
		return nil, err
	}

	return list, nil
}
