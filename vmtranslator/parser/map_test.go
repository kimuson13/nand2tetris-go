package parser

import (
	"testing"
	"vmtranslator/codewriter"
)

func TestMapCommandTypeToArithmeticKind(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      commandType
		want    codewriter.ArithmeticKind
		wantErr bool
	}{
		"add":  {ADD, codewriter.ADD, noErr},
		"hoge": {commandType("hoge"), 0, wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := mapCommandTypeToArithmeticKind(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("no err: %v", tc.in)
			}

			if got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}
