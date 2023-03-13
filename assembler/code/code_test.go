package code_test

import (
	"assembler/code"
	"testing"
)

func TestCCommand_Convert(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		dest    *code.Dest
		comp    code.Comp
		jump    *code.Jump
		want    string
		wantErr bool
	}{
		"D=M":      {ptr(code.DEST_D), code.COMP_M, nil, "1111110000010000", noErr},
		"D=D-A":    {ptr(code.DEST_D), code.COMP_D_MINUS_A, nil, "1110010011010000", noErr},
		"D=D-M":    {ptr(code.DEST_D), code.COMP_D_MINUS_M, nil, "1111010011010000", noErr},
		"D=M;JMP":  {ptr(code.DEST_D), code.COMP_M, ptr(code.JUMP), "1111110000010111", noErr},
		"D;JGT":    {nil, code.COMP_D, ptr(code.JUMP_GREATER_THAN), "1110001100000001", noErr},
		"hoge=M+1": {ptr(code.Dest("hoge")), code.COMP_M_ADD_1, nil, "", wantErr},
		"M=hoge":   {ptr(code.DEST_M), code.Comp("hoge"), nil, "", wantErr},
		"M;hoge":   {nil, code.COMP_M, ptr(code.Jump("hoge")), "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cCommand := code.CCommand{
				Dest: tc.dest,
				Comp: tc.comp,
				Jump: tc.jump,
			}

			got, err := cCommand.Convert()
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Error("no error")
			}

			if got != tc.want {
				t.Errorf("got = %v, but got = %v", tc.want, got)
			}
		})
	}
}

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
				Symbol:  tc.symbol,
				Address: tc.value,
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

func ptr[T any](val T) *T {
	return &val
}
