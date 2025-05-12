package admin

import "context"

type Service struct {
	repo      Repository
	validator *Validator
}

func NewAdminService(repo Repository, validator *Validator) *Service {
	return &Service{repo: repo, validator: validator}
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*User, error) {
	return s.repo.GetAllUsers(ctx)
}
