package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	GetUserByID(ctx context.Context, id string) (*Profile, error)
	UpdateUserProfile(ctx context.Context, id string, p *Profile) (*Profile, error)
}

type SQLRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	const query = `
		INSERT INTO users (id, phone, first_name, last_name, password, is_deleted)
		VALUES ($1, $2, $3, $4, $5, $6)
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
		user.IsDeleted,
	).Scan(&user.CreatedAt)

	if err != nil {
		if isDuplicateError(err) {
			return nil, ErrPhoneExists
		}

		return nil, fmt.Errorf("user repository create: %w", err)
	}

	const criteriaQuery = `
		INSERT INTO user_criteria (user_id)
		VALUES ($1)
	`

	_, err = r.db.ExecContext(
		ctx,
		criteriaQuery,
		user.ID,
	)

	if err != nil {
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
		return nil, fmt.Errorf("get user by phone: %w", err)
	}

	return &user, nil
}

func (r *SQLRepository) GetUserByID(ctx context.Context, id string) (*Profile, error) {
	const query = `
		SELECT phone, first_name, last_name
		FROM users
		WHERE id = $1
	`

	var profile Profile
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&profile.Phone,
		&profile.FirstName,
		&profile.LastName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrInvalidCredentials
	}

	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	const critQuery = `
		SELECT age, job, gender, is_married
		FROM user_criteria
		WHERE user_id = $1
	`

	err = r.db.QueryRowContext(ctx, critQuery, id).Scan(
		&profile.Age,
		&profile.Job,
		&profile.Gender,
		&profile.IsMarried,
	)

	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return &profile, nil
}

func isDuplicateError(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23505"
}

func (r *SQLRepository) UpdateUserProfile(ctx context.Context, id string, p *Profile) (*Profile, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	if p.FirstName != nil || p.LastName != nil {
		query := `UPDATE users SET`
		args := []interface{}{}
		idx := 1

		if p.FirstName != nil {
			query += fmt.Sprintf(" first_name = $%d,", idx)
			args = append(args, *p.FirstName)
			idx++
		}

		if p.LastName != nil {
			query += fmt.Sprintf(" last_name = $%d,", idx)
			args = append(args, *p.FirstName)
			idx++
		}

		query = strings.TrimSuffix(query, ",") + fmt.Sprintf(" WHERE id = $%d", idx)
		args = append(args, id)

		if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
			return nil, fmt.Errorf("update users: %w", err)
		}
	}

	if p.Age != nil || p.Job != nil || p.Gender != nil || p.IsMarried != nil {
		query := `UPDATE user_criteria SET`
		args := []interface{}{}
		idx := 1

		if p.Age != nil {
			query += fmt.Sprintf(" age = $%d,", idx)
			args = append(args, *p.Age)
			idx++
		}

		if p.Job != nil {
			query += fmt.Sprintf(" job = $%d,", idx)
			args = append(args, *p.Job)
			idx++
		}

		if p.Gender != nil {
			query += fmt.Sprintf(" gender = $%d,", idx)
			args = append(args, *p.Gender)
			idx++
		}

		if p.IsMarried != nil {
			query += fmt.Sprintf(" is_married = $%d,", idx)
			args = append(args, *p.IsMarried)
			idx++
		}

		query = strings.TrimSuffix(query, ",") + fmt.Sprintf(" WHERE user_id = $%d", idx)
		args = append(args, id)

		if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
			return nil, fmt.Errorf("update users: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return r.GetUserByID(ctx, id)
}
