package repository

import (
	"context"
	"financial/domain/entity"
)

type AccountRepository interface {
	FindByClientIdAndCourseId(ctx context.Context, clientID, courseID string) (*entity.Account, error)
	Save(ctx context.Context, account *entity.Account) (string, error)
}
