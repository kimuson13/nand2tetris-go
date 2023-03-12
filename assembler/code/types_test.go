package code

import "testing"

func TestMapDestToBinary(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      Dest
		want    string
		wantErr bool
	}{
		"ok_m":    {DEST_M, "001", noErr},
		"ng_hoge": {Dest("hoge"), "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := mapDestToBinary(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Error("error not occured")
			}

			if got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}

func TestMapCompToBinary(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      Comp
		want    string
		wantErr bool
	}{
		"ok_comp_not_M": {COMP_NOT_M, "1110001", noErr},
		"ng_hoge":       {Comp("hoge"), "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := mapCompToBinary(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Error("error not occured")
			}

			if got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}

func TestMapJumpToBinary(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      Jump
		want    string
		wantErr bool
	}{
		"ok_jmp":  {JUMP, "111", noErr},
		"ng_hoge": {Jump("hoge"), "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := mapJumpToBinary(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Error("error not occured")
			}

			if got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}
