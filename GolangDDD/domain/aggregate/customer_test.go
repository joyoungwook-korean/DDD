package aggregate_test

import (
	"testing"

	"github.com/percybolmer/ddd-go/domain/aggregate"
)

func TestNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	tests := []testCase{
		{
			test:        "Empty name",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "valid",
			name:        "Young Wook jo",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tt.name)
			if err != tt.expectedErr {
				t.Errorf("Expected Error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
