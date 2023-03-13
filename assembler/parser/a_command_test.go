package parser

import (
	"testing"
)

func TestIsACommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want bool
	}{
		"no_symbol":            {"@123", true},
		"with_symbol":          {"@hoge", true},
		"include_underscore":   {"@HOGE_HOGE", true},
		"not_start_@":          {"hoge", false},
		"start_num_symbol":     {"@1hoge", false},
		"include_invalid_char": {"@hoge,", false},
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
