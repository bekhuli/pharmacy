package admin

import "context"

type Service struct {
	repo      Repository
	validator *Validator
}

func NewAdminService(repo Repository, validator *Validator) *Service {
	return &Service{repo: repo, validator: validator}
}

func (s *Service) GetAllUsers(ctx context.Context, page, limit int) ([]*User, int, error) {
	offset := (page - 1) * limit
	return s.repo.GetAllUsers(ctx, offset, limit)
}
