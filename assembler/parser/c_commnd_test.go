package parser

import "testing"

func TestIsCCommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want bool
	}{
		"ok_only_eq":            {"hoge=huga", true},
		"ok_only_semi":          {"hoge;huga", true},
		"ok_full":               {"hoge=huga;piyo", true},
		"ng_too_many_eq":        {"hoge=huga==", false},
		"ng_too_many_semi":      {"hoge;;;;huga", false},
		"ng_switch_eq_and_semi": {"hoge;huga=piyo", false},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := isCComand(tc.raw); got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}
