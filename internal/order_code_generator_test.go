package internal

import "testing"

func TestOrderTypeCodeGenerator(t *testing.T) {
	var tests = map[string]struct {
		orderType string
		expected  string
	}{
		"food": {
			orderType: "food",
			expected:  "fo-12313-200911",
		},
		"din-in": {
			orderType: "pro",
			expected:  "pro-12313-200911",
		},
		"subscription": {
			orderType: "subscription",
			expected:  "sub-12313-200911",
		},
	}

	cg := NewCodeGenerator()
	for _, test := range tests {
		code, err := cg.GetOrderCode(test.orderType)

		if err != nil {
			t.Failed()
		}

		if code != test.expected {
			t.Failed()
		}
	}
}
