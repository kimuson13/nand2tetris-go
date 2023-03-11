package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

func TestGenCComandStmt(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    CCommandStmt
		wantErr bool
	}{
		"dest_comp_jump": {"hoge=huga;piyo", ccstmt(csRaw("hoge=huga;piyo"), eqPos(4), semiCPos(9), kind(DEST_COMP_JUMP)), noErr},
		"dest_comp":      {"hoge=huga", ccstmt(csRaw("hoge=huga"), eqPos(4), kind(DEST_COMP)), noErr},
		"comp_jump":      {"hoge;JMP", ccstmt(csRaw("hoge;JMP"), semiCPos(4), kind(COMP_JUMP)), noErr},
		"no_much_error":  {"hoge", ccstmt(csRaw("hoge")), wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := genCCommandStmt(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(CCommandStmt{})); diff != "" {
				t.Errorf("want = %#v\ngot =%#v\ndiff=%s", tc.want, got, diff)
			}
		})
	}
}

type CComandStmtOption = Option[*CCommandStmt]

func ccstmt(opts ...CComandStmtOption) CCommandStmt {
	ccstmt := CCommandStmt{eqPos: -1, semiColonPos: -1}

	for _, opt := range opts {
		opt(&ccstmt)
	}

	return ccstmt
}

func csRaw(v string) CComandStmtOption {
	return func(val *CCommandStmt) {
		val.raw = v
	}
}

func eqPos(v int) CComandStmtOption {
	return func(val *CCommandStmt) {
		val.eqPos = v
	}
}

func semiCPos(v int) CComandStmtOption {
	return func(val *CCommandStmt) {
		val.semiColonPos = v
	}
}

func kind(v CCommandKind) CComandStmtOption {
	return func(val *CCommandStmt) {
		val.kind = v
	}
}
