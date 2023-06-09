package parser

import (
	"assembler/code"
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
			if got := isCCommand(tc.raw); got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}

func TestParse(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      cCommand
		want    *code.CCommand
		wantErr bool
	}{
		"ok_full": {
			in:      *ccmd(pDest("M"), pComp("0"), pJump("JMP")),
			want:    codeCmd(cDest(code.DEST_M), cComp(code.COMP_0), cJump(code.JUMP)),
			wantErr: noErr,
		},
		"ok_dest_comp": {
			in:      *ccmd(pDest("AM"), pComp("1")),
			want:    codeCmd(cDest(code.DEST_AM), cComp(code.COMP_1)),
			wantErr: noErr,
		},
		"ok_comp_jump": {
			in:      *ccmd(pComp("A"), pJump("JMP")),
			want:    codeCmd(cComp(code.COMP_A), cJump(code.JUMP)),
			wantErr: noErr,
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := tc.in.parse()
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(cCommand{})); diff != "" {
				t.Errorf("diff=%v", diff)
			}
		})
	}
}

func TestMapCodeDest(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    *code.Dest
		wantErr bool
	}{
		"ok_M":    {"M", ptr(code.DEST_M), noErr},
		"ok_AMD":  {"AMD", ptr(code.DEST_AMD), noErr},
		"ng_hoge": {"hoge", nil, wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := mapCodeDest(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if got != nil && *got != *tc.want {
				t.Errorf("want = %v, but got = %v", *tc.want, *got)
			}
		})
	}
}

func TestMapCodeComp(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    code.Comp
		wantErr bool
	}{
		"ok_0":    {"0", code.COMP_0, noErr},
		"ok_-A":   {"-A", code.COMP_MINUS_A, noErr},
		"ok_!D":   {"!D", code.COMP_NOT_D, noErr},
		"ok_A+1":  {"A+1", code.COMP_A_ADD_1, noErr},
		"ok_D|A":  {"D|A", code.COMP_D_OR_A, noErr},
		"ok_D-M":  {"D-M", code.COMP_D_MINUS_M, noErr},
		"ok_D&M":  {"D&M", code.COMP_D_AND_M, noErr},
		"ng_hoge": {"hoge", "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := mapCodeComp(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}

func TestMapCodeJump(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    *code.Jump
		wantErr bool
	}{
		"ok_JGT":  {"JGT", ptr(code.JUMP_GREATER_THAN), noErr},
		"ok_JMP":  {"JMP", ptr(code.JUMP), noErr},
		"ng_hoge": {"hoge", nil, wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := mapCodeJump(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if got != nil && *got != *tc.want {
				t.Errorf("want = %v, but got = %v", *tc.want, *got)
			}
		})
	}
}

func TestGenCComandStmt(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    cCommandStmt
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

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(cCommandStmt{})); diff != "" {
				t.Errorf("want = %#v\ngot =%#v\ndiff=%s", tc.want, got, diff)
			}
		})
	}
}

func TestStmtToCCommand(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want *cCommand
	}{
		"dest_comp_jump": {"hoge=huga;piyo", ccmd(pDest("hoge"), pComp("huga"), pJump("piyo"))},
		"dest_comp":      {"hoge=huga", ccmd(pDest("hoge"), pComp("huga"))},
		"comp_jump":      {"huga;piyo", ccmd(pComp("huga"), pJump("piyo"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stmt, err := genCCommandStmt(tc.in)
			if err != nil {
				t.Fatal(err)
			}

			got, err := stmt.toCCommand()
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(cCommand{})); diff != "" {
				t.Errorf("want != got\ndiff=%s", diff)
			}
		})
	}
}

func TestToDestCompJump(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want *cCommand
	}{
		"normal": {"hoge=huga;piyo", ccmd(pDest("hoge"), pComp("huga"), pJump("piyo"))},
		"long":   {"hogehogehoge=hugahugahhh;pipipioooo", ccmd(pDest("hogehogehoge"), pComp("hugahugahhh"), pJump("pipipioooo"))},
		"short":  {"a=b;c", ccmd(pDest("a"), pComp("b"), pJump("c"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stmt, err := genCCommandStmt(tc.in)
			if err != nil {
				t.Fatal(err)
			}

			got, err := stmt.toDestCompJump()
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(cCommand{})); diff != "" {
				t.Errorf("want != got\ndiff=%s", diff)
			}
		})
	}
}

func TestToDestComp(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want *cCommand
	}{
		"normal": {"hoge=huga", ccmd(pDest("hoge"), pComp("huga"))},
		"long":   {"hogehogehoge=hugahugahhh", ccmd(pDest("hogehogehoge"), pComp("hugahugahhh"))},
		"short":  {"a=b", ccmd(pDest("a"), pComp("b"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stmt, err := genCCommandStmt(tc.in)
			if err != nil {
				t.Fatal(err)
			}

			got, err := stmt.toDestComp()
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(cCommand{})); diff != "" {
				t.Errorf("want != got\ndiff=%s", diff)
			}
		})
	}
}

func TestToCompJump(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want *cCommand
	}{
		"normal": {"huga;piyo", ccmd(pComp("huga"), pJump("piyo"))},
		"long":   {"hugahugahhh;pipipioooo", ccmd(pComp("hugahugahhh"), pJump("pipipioooo"))},
		"short":  {"b;c", ccmd(pComp("b"), pJump("c"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			stmt, err := genCCommandStmt(tc.in)
			if err != nil {
				t.Fatal(err)
			}

			got, err := stmt.toCompJump()
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(cCommand{})); diff != "" {
				t.Errorf("want != got\ndiff=%s", diff)
			}
		})
	}
}

func ptr[T any](val T) *T {
	return &val
}

type CComandStmtOption = Option[*cCommandStmt]
type CCommandOption = Option[*cCommand]
type CodeCComandOption = Option[*code.CCommand]

func ccstmt(opts ...CComandStmtOption) cCommandStmt {
	ccstmt := cCommandStmt{eqPos: -1, semiColonPos: -1}

	for _, opt := range opts {
		opt(&ccstmt)
	}

	return ccstmt
}

func ccmd(opts ...CCommandOption) *cCommand {
	cCommand := &cCommand{}

	for _, opt := range opts {
		opt(cCommand)
	}

	return cCommand
}

func codeCmd(opts ...CodeCComandOption) *code.CCommand {
	codeCComand := &code.CCommand{
		Dest: nil,
		Comp: "",
		Jump: nil,
	}

	for _, opt := range opts {
		opt(codeCComand)
	}

	return codeCComand
}

func cDest(v code.Dest) CodeCComandOption {
	return func(val *code.CCommand) {
		val.Dest = &v
	}
}

func cComp(v code.Comp) CodeCComandOption {
	return func(val *code.CCommand) {
		val.Comp = v
	}
}

func cJump(v code.Jump) CodeCComandOption {
	return func(val *code.CCommand) {
		val.Jump = &v
	}
}

func csRaw(v string) CComandStmtOption {
	return func(val *cCommandStmt) {
		val.raw = v
	}
}

func eqPos(v int) CComandStmtOption {
	return func(val *cCommandStmt) {
		val.eqPos = v
	}
}

func semiCPos(v int) CComandStmtOption {
	return func(val *cCommandStmt) {
		val.semiColonPos = v
	}
}

func kind(v cCommandKind) CComandStmtOption {
	return func(val *cCommandStmt) {
		val.kind = v
	}
}

func pDest(v string) CCommandOption {
	return func(val *cCommand) {
		val.dest = v
	}
}

func pComp(v string) CCommandOption {
	return func(val *cCommand) {
		val.comp = v
	}
}

func pJump(v string) CCommandOption {
	return func(val *cCommand) {
		val.jump = v
	}
}
