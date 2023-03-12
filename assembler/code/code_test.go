package code_test

import (
	"assembler/code"
	"testing"
)

func TestACommand_Convert(t *testing.T) {
	testCases := map[string]struct {
		symbol string
		value  int
		want   string
	}{
		"@100": {"", 100, "0000000001100100"},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			aCommand := code.ACommand{
				Symbol: tc.symbol,
				Value:  tc.value,
			}

			got, err := aCommand.Convert()
			if err != nil {
				t.Error(err)
			}

			if got != tc.want {
				t.Errorf("want:\n%v\ngot:\n%v", tc.want, got)
			}
		})
	}
}
