package admin

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAllUsers(ctx context.Context, offset, limit int) ([]*User, int, error)
	GetUser(ctx context.Context, userID string) (*User, error)
}

type SQLRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) GetAllUsers(ctx context.Context, offset, limit int) ([]*User, int, error) {
	const usersQuery = `
		SELECT id, phone, first_name, last_name, created_at
		FROM users
		WHERE is_deleted = false
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, usersQuery, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get all users: %w", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Phone, &u.FirstName, &u.LastName, &u.CreatedAt)
		if err != nil {
			return nil, 0, fmt.Errorf("scan user: %w", err)
		}

		var roleID string
		err = r.db.QueryRowContext(ctx, `SELECT role_id FROM user_roles WHERE user_id = $1`, u.ID).Scan(&roleID)
		if err != nil {
			return nil, 0, fmt.Errorf("error getting roleID: %w", err)
		}

		err = r.db.QueryRowContext(ctx, `SELECT name FROM roles WHERE id = $1`, roleID).Scan(&u.Role)
		if err != nil {
			return nil, 0, fmt.Errorf("error getting role: %w", err)
		}

		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows err: %w", err)
	}

	const countQuery = `
		SELECT COUNT(*) FROM users WHERE is_deleted = false
	`

	var total int
	err = r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count users: %w", err)
	}

	return users, total, nil
}

func (r *SQLRepository) GetUser(ctx context.Context, userID string) (*User, error) {
	const userQuery = `
		SELECT id, phone, first_name, last_name, created_at
		FROM users
		WHERE id = $1
	`

	var user User
	err := r.db.QueryRowContext(ctx, userQuery, userID).Scan(
		&user.ID,
		&user.Phone,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	var roleID string
	err = r.db.QueryRowContext(ctx, `SELECT role_id FROM user_roles WHERE user_id = $1`, userID).Scan(&roleID)
	if err != nil {
		return nil, fmt.Errorf("error getting role: %w", err)
	}

	err = r.db.QueryRowContext(ctx, `SELECT name FROM roles WHERE id = $1`, roleID).Scan(&user.Role)
	if err != nil {
		return nil, fmt.Errorf("error getting role: %w", err)
	}

	return &user, nil
}
