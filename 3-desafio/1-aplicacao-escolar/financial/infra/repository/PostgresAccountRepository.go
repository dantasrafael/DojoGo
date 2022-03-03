package repository

import (
	"context"
	"database/sql"
	"financial/domain/entity"
	"financial/domain/repository"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

const (
	selectAccountByClientAndCourse = "SELECT id, client_id, course_id, installments, total, created_at FROM account WHERE client_id = $1 AND course_id = $2 "
	insertAccount                  = "INSERT INTO account (client_id, course_id, installments, total) VALUES ( $1, $2, $3, $4) RETURNING id"
)

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository(db *sql.DB) repository.AccountRepository {
	return PostgresAccountRepository{db: db}
}

func (p PostgresAccountRepository) FindByClientIdAndCourseId(ctx context.Context, clientID, courseID string) (*entity.Account, error) {
	var account entity.Account
	err := p.db.QueryRowContext(ctx, selectAccountByClientAndCourse, clientID, courseID).
		Scan(&account.ID, &account.ClientID, &account.CourseID, &account.Installments, &account.Total, &account.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not find account by client[%s] and course[%s]: %v", clientID, courseID, err)
	}
	return &account, nil
}

func (p PostgresAccountRepository) Save(ctx context.Context, account *entity.Account) (*uuid.UUID, error) {
	var idStr string
	err := p.db.QueryRowContext(ctx, insertAccount, account.ClientID, account.CourseID, account.Installments, account.Total).
		Scan(&idStr)
	if err != nil {
		return nil, fmt.Errorf("could not save account with client[%s] and course[%s]: %v", account.ClientID, account.CourseID, err)
	}
	if id, err := uuid.FromString(idStr); err != nil {
		return nil, fmt.Errorf("could not parse id[%s] to uuid: %v", idStr, err)
	} else {
		return &id, nil
	}
}
