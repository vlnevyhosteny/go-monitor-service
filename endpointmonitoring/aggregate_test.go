package endpointmonitoring_test

import (
	"testing"

	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
	"github.com/vlnevyhosteny/go-monitor-service/endpointmonitoring"
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
			expectedErr: endpointmonitoring.ErrEmptyEndpointName,
		},
	}

	for name, test := range newCustomerTests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var _, err = endpointmonitoring.NewEndpointMonitoring(test.name, test.address)

			if err != test.expectedErr {
				t.Fatalf("Expected error %v, got %v", test.expectedErr, err)
			}
		})
	}
}
