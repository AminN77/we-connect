package internal

import "context"

// Service is the aggregator of the internal(domain) layer
type Service interface {
	Insert(fd *FinancialData) error
	Get(q *Query, ctx context.Context) ([]*FinancialData, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Insert(fd *FinancialData) error {
	return s.repo.Insert(fd)
}

func (s *service) Get(q *Query, ctx context.Context) ([]*FinancialData, error) {
	return s.repo.Get(q, ctx)
}
