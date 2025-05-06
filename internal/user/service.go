package user

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      Repository
	validator *Validator
}

func NewUserService(repo Repository, validator *Validator) *Service {
	return &Service{repo: repo, validator: validator}
}

func (s *Service) RegisterUser(ctx context.Context, dto RegisterRequest) (*User, error) {
	if err := s.validator.Validate(dto); err != nil {
		return nil, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user := &User{
		ID:        uuid.New(),
		Phone:     dto.Phone,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Password:  string(hashedPassword),
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *Service) LoginUser(ctx context.Context, dto LoginRequest) (*User, error) {
	if err := s.validator.Validate(dto); err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByPhone(ctx, dto.Phone)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
