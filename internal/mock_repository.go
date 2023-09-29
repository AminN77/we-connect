package internal

import (
	"context"
	"errors"
)

var (
	ErrMock = errors.New("some error")
)

type mockRepository struct {
}

func (*mockRepository) Insert(fd *FinancialData) error {
	return nil
}

func (*mockRepository) InsertBatch(fd []*FinancialData) error {
	return nil
}

func (*mockRepository) Get(q *Query, ctx context.Context) ([]*FinancialData, error) {
	if q == nil {
		return nil, ErrMock
	}
	return []*FinancialData{}, nil
}
