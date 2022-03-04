package usecase

import (
	"folha/src/domain/entities"
	repository "folha/src/infra/repositories"
)

func EmployeeCreate(model *entities.Employee) error {
	if err := repository.EmployeeSave(model); err != nil {
		return err
	}

	return nil
}
