package admin

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
}

type SQLRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) GetAllUsers(ctx context.Context) ([]*User, error) {
	const usersQuery = `
		SELECT id, phone, first_name, last_name, created_at
		FROM users
		WHERE is_deleted = false
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, usersQuery)
	if err != nil {
		return nil, fmt.Errorf("get all users: %w", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Phone, &u.FirstName, &u.LastName, &u.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return users, nil
}
