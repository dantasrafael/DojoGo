package repository

import "financial/domain/entity"

type AccountRepository interface {
	FindByStudentIdAndCourseId(studentID, courseID string) (*entity.Account, error)
	Save(account *entity.Account) (string, error)
}
