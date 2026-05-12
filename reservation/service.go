package reservation

import "context"

type Repository interface {
	List(ctx context.Context) ([]Reservation, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]Reservation, error) {
	return s.repo.List(ctx)
}
