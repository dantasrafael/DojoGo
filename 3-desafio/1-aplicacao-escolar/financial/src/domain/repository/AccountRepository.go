package repository

import (
	"context"
	"financial/src/domain/entity"
	"github.com/satori/go.uuid"
)

type AccountRepository interface {
	FindByClientIdAndCourseId(ctx context.Context, clientID, courseID string) (*entity.Account, error)
	Save(ctx context.Context, account *entity.Account) (*uuid.UUID, error)
}
