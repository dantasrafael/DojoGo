package usecase

import (
	"folha/src/domain/entities"
	"folha/src/infra/repositories"
)

func EmployeeFindAll() ([]*entities.Employee, error) {
	return repositories.EmployeeFindAll()
}

func EmployeeFindById(id *uint64) (*entities.Employee, error) {
	return repositories.EmployeeFindById(id);
}