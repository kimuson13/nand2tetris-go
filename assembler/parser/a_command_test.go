package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestToACommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want *aCommand
	}{
		"ok_with_no_symbol": {"@123", a(aAddress(123))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := toACommand(tc.raw)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(aCommand{})); diff != "" {
				t.Errorf("want = %#v, got = %#v, \ndiff: %s", tc.want, got, diff)
			}
		})
	}
}

type ACommandOption func(val *aCommand)

func a(opts ...ACommandOption) *aCommand {
	aCommand := &aCommand{
		address: 0,
		symbol:  "",
	}

	for _, opt := range opts {
		opt(aCommand)
	}

	return aCommand
}

func aAddress(v int) ACommandOption {
	return func(val *aCommand) {
		val.address = v
	}
}

func aSymbol(v string) ACommandOption {
	return func(val *aCommand) {
		val.symbol = v
	}
}
