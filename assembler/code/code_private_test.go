package code

import (
	"testing"
)

func TestCCommand_dest(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      *Dest
		want    string
		wantErr bool
	}{
		"ok_D":    {ptr(DEST_D), "010", noErr},
		"ok_nil":  {nil, "000", noErr},
		"ng_hoge": {ptr(Dest("hoge")), "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cCommand := CCommand{
				Dest: tc.in,
			}

			got, err := cCommand.dest()
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Error("no err")
			}

			if got != tc.want {
				t.Errorf("want = %v, got = %v", tc.want, got)
			}
		})
	}
}

func TestCCommand_comp(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      Comp
		want    string
		wantErr bool
	}{
		"ok_COMP_D_OR_A": {COMP_D_OR_A, "0010101", noErr},
		"ng_hoge":        {Comp("hoge"), "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cCommand := CCommand{
				Comp: tc.in,
			}

			got, err := cCommand.comp()
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Error("no err")
			}

			if got != tc.want {
				t.Errorf("want = %v, got = %v", tc.want, got)
			}
		})
	}
}

func TestCCommand_jump(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      *Jump
		want    string
		wantErr bool
	}{
		"ok_JLT":  {ptr(JUMP_LOWER_THAN), "100", noErr},
		"ok_nil":  {nil, "000", noErr},
		"ng_hoge": {ptr(Jump("hoge")), "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cCommand := CCommand{
				Jump: tc.in,
			}

			got, err := cCommand.jump()
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Error("no err")
			}

			if got != tc.want {
				t.Errorf("want = %v, got = %v", tc.want, got)
			}
		})
	}
}

func ptr[T any](val T) *T {
	return &val
}
