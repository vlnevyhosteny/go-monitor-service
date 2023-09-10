package aggregates_test

import (
	"testing"

	"github.com/vlnevyhosteny/go-monitor-service/domain/aggregates"
	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

func TestNewEndpointMonitoring(t *testing.T) {
	var address, _ = valueobjects.NewAddress("http://test.com")

	var newCustomerTests = map[string]struct {
		name        string
		address     valueobjects.Address
		result      string
		expectedErr error
	}{
		"success": {
			name:        "test",
			address:     address,
			result:      "test",
			expectedErr: nil,
		},
		"empty name": {
			name:        "",
			address:     address,
			result:      "",
			expectedErr: aggregates.ErrEmptyEndpointName,
		},
	}

	for name, test := range newCustomerTests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var _, err = aggregates.NewEndpointMonitoring(test.name, test.address)

			if err != test.expectedErr {
				t.Fatalf("Expected error %v, got %v", test.expectedErr, err)
			}
		})
	}
}
