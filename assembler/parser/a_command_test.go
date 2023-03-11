package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToACommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want ACommand
	}{}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := toACommand(tc.raw)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(ACommand{})); diff != "" {
				t.Errorf("want = %#v, got = %#v, \ndiff: %s", tc.want, got, diff)
			}
		})
	}
}

type ACommandOption func(val *ACommand)

func a(opts ...ACommandOption) ACommand {
	aCommand := ACommand{
		value:  0,
		symbol: "",
	}

	for _, opt := range opts {
		opt(&aCommand)
	}

	return aCommand
}
