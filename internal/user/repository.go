package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
}

type SQLRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	const query = `
		INSERT INTO users (id, phone, first_name, last_name, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.ID,
		user.Phone,
		user.FirstName,
		user.LastName,
		user.Password,
	).Scan(&user.CreatedAt)

	if err != nil {
		if isDuplicateError(err) {
			return nil, ErrPhoneExists
		}

		return nil, fmt.Errorf("user repository create: %w", err)
	}

	return user, nil
}

func (r *SQLRepository) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	const query = `
		SELECT id, phone, first_name, last_name, password, created_at
		FROM users
		WHERE phone = $1
	`

	var user User
	err := r.db.QueryRowContext(ctx, query, phone).Scan(
		&user.ID,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrInvalidCredentials
	}

	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	return &user, nil
}

func isDuplicateError(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23505"
}
