package parser

import (
	"testing"
)

func TestIsACommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want bool
	}{
		"ok_no_symbol":   {"@123", true},
		"ok_with_symbol": {"@hoge", true},
		"ng_not_start_@": {"hoge", false},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := isACommand(tc.raw); got != tc.want {
				t.Errorf("want = %v, got = %v", tc.want, got)
			}
		})
	}
}
