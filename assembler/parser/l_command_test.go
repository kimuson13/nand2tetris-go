package parser

import (
	"testing"
)

func TestIsLCommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want bool
	}{
		"ok":            {"(hoge)", true},
		"ng_no_()":      {"hoge", false},
		"ng_head_not_(": {"hoge)", false},
		"ng_tail_not_)": {"(hoge", false},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := isLCommand(tc.raw); got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}
