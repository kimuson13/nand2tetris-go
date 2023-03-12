package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestToLCommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want *LCommand
	}{
		"ok": {"(hoge)", l(lSymbol("hoge"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := toLCommand(tc.raw)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(LCommand{})); diff != "" {
				t.Errorf("want = %#v, but got = %#v\ndiff=%v", tc.want, got, diff)
			}
		})
	}
}

type LCommandOption Option[*LCommand]

func l(opts ...LCommandOption) *LCommand {
	lCommand := &LCommand{
		symbol: "",
	}

	for _, opt := range opts {
		opt(lCommand)
	}

	return lCommand
}

func lSymbol(v string) LCommandOption {
	return func(val *LCommand) {
		val.symbol = v
	}
}
