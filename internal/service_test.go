package internal

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceGet(t *testing.T) {
	testCases := []struct {
		name        string
		q           *Query
		srv         Service
		expected    []*FinancialData
		expectedErr error
	}{
		{
			name:        "success",
			q:           &Query{},
			srv:         NewService(&mockRepository{}),
			expected:    []*FinancialData{},
			expectedErr: nil,
		},
		{
			name:        "query nil & repo err",
			q:           nil,
			srv:         NewService(&mockRepository{}),
			expected:    nil,
			expectedErr: ErrMock,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.srv.Get(tc.q, context.Background())
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestServiceInsert(t *testing.T) {
	testCases := []struct {
		name        string
		srv         Service
		expectedErr error
	}{
		{
			name:        "success",
			srv:         NewService(&mockRepository{}),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.srv.Insert(&FinancialData{})
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
