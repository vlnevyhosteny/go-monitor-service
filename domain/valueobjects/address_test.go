package valueobjects_test

import (
	"strings"
	"testing"

	"github.com/vlnevyhosteny/go-monitor-service/domain/valueobjects"
)

func TestNewAddress(t *testing.T) {
	var newAddressTests = map[string]struct {
		input       string
		expectedErr error
		expectedUrl string
	}{
		"success": {
			input:       "http://test.com",
			expectedErr: nil,
			expectedUrl: "http://test.com",
		},
		"empty input": {
			input:       "",
			expectedErr: valueobjects.ErrEmptyInputUrl,
			expectedUrl: "",
		},
		"invalid input": {
			input:       "hp/there?",
			expectedErr: valueobjects.ErrInvalidInputUrl(nil),
			expectedUrl: "",
		},
	}

	for name, test := range newAddressTests {
		test := test

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var address, err = valueobjects.NewAddress(test.input)

			if test.expectedErr != nil && strings.Contains(err.Error(), test.expectedErr.Error()) == false {
				t.Fatalf("Expected error %v, got %v", test.expectedErr, err)
			}

			if address.Url.String() != test.expectedUrl {
				t.Fatalf("Expected url %v, got %v", test.expectedUrl, address.Url.String())
			}
		})
	}
}
