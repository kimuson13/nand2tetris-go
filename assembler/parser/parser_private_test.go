package parser

import (
	"assembler/symtable"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSymbolicLink(t *testing.T) {
	b := "// comment\r\n@R0\r\n@123\r\n(HOGE)\r\n@YEAR\r\n@HOGE"
	parser, close := setUp(t, []byte(b))
	wants := []struct {
		symbol  string
		address int
	}{
		{"HOGE", 2},
		{"YEAR", 16},
	}

	if err := parser.SymbolicLink(); err != nil {
		t.Error(err)
	}

	for _, want := range wants {
		got, err := parser.symbolTable.GetAddress(want.symbol)
		if err != nil {
			t.Error(err)
		}

		if got != want.address {
			t.Errorf("want = %d, but got = %d", want.address, got)
		}
	}

	close()
}

func TestGetCommands(t *testing.T) {
	t.Parallel()
	in := s("@123", "// comment", "(hoge) // hello", "", "M=M+1;JMP")
	want := s("@123", "(hoge)", "M=M+1;JMP")

	got, err := getCommands(in)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("want:\n%v\ngot:\n%v\ndiff:\n%s", want, got, diff)
	}
}

func TestGetCommand(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"ok_a_command":        {"@123", "@123", noErr},
		"ok_a_command_tab":    {"	@123", "@123", noErr},
		"ok_skip_comment":     {"//comment", "", noErr},
		"ok_empty":            {"", "", noErr},
		"ng_between_space":    {"@12 3", "", wantErr},
		"ng_too_many_comment": {"@123 // hoge //huga", "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := getCommand(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if got != tc.want {
				t.Errorf("want = %s, but got = %s", tc.want, got)
			}
		})
	}
}

func TestLinkLCommandSymbol(t *testing.T) {
	wants := []struct {
		symbol  string
		address int
	}{
		{"HOGE", 2},
		{"HUGA", 3},
	}

	b := []byte("//comment\r\n@R1\r\n@123\r\n(HOGE)\r\n\r\nM=M+1\r\n(HUGA) // hoge")
	parser, close := setUp(t, b)

	if err := parser.linkLCommandSymbol(); err != nil {
		t.Error(err)
	}

	for _, want := range wants {
		got, err := parser.symbolTable.GetAddress(want.symbol)
		if err != nil {
			t.Error(err)
		}

		if got != want.address {
			t.Errorf("want = %d, but got = %d", want.address, got)
		}
	}

	close()
}

func TestLinkACommandSymbol(t *testing.T) {
	wants := []struct {
		symbol  string
		address int
	}{
		{"R1", 1},
		{"YEAR", 16},
		{"HOO", 17},
	}

	b := []byte("//comment\r\n@R1\r\n@123\r\n(HOGE)\r\n\r\nM=M+1\r\n(HUGA) // hoge\r\n@YEAR\r\n@HOO")
	parser, close := setUp(t, b)

	if err := parser.linkACommandSymbol(); err != nil {
		t.Error(err)
	}

	for _, want := range wants {
		got, err := parser.symbolTable.GetAddress(want.symbol)
		if err != nil {
			t.Error(err)
		}

		if got != want.address {
			t.Errorf("want = %d, but got = %d", want.address, got)
		}
	}

	close()
}

func setUp(t *testing.T, body []byte) (Parser, func()) {
	t.Helper()
	f, err := os.CreateTemp("./", "")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := f.Write(body); err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}

	parser, err := New(f.Name())
	if err != nil {
		os.Remove(f.Name())
		t.Fatal(err)
	}
	parser.symbolTable = symtable.New()

	return parser, func() { os.Remove(f.Name()) }
}

func TestAddEntryToSymTable(t *testing.T) {
	parser := p()
	in := struct {
		symbol  string
		address int
	}{
		"hoge",
		123,
	}
	want := 123

	if err := parser.addEntryToSymTable(in.symbol, in.address); err != nil {
		t.Fatal(err)
	}

	got, err := parser.symbolTable.GetAddress(in.symbol)
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("want = %d, but got = %d", want, got)
	}
}

func TestHasMoreCommand(t *testing.T) {
	testCases := map[string]struct {
		commands   []string
		currentIdx int
		want       bool
	}{
		"hasMoreCommand": {s("h", "o"), 1, true},
		"noMoreCommand":  {s("h", "0"), 2, false},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			parser := p(commands(tc.commands...), currentIdx(tc.currentIdx))
			got := parser.hasMoreCommand()
			if got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}

func TestAdvance(t *testing.T) {
	parser := p(commands("hoge", "huga"))

	want := 2
	cnt := 0
	for parser.hasMoreCommand() {
		cnt++
		parser.advance()
	}

	if cnt != want {
		t.Errorf("want = %d, but got = %d", want, cnt)
	}
}

func TestCommandType(t *testing.T) {
	testCases := map[string]struct {
		in   []string
		want Command
	}{
		"a_command": {s("@123"), a(aAddress(123))},
		"l_command": {s("(hoge)"), l(lSymbol("hoge"))},
		"c_command": {s("hoge=huga;piyo"), ccmd(pDest("hoge"), pComp("huga"), pJump("piyo"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			parser := p(commands(tc.in...))
			got, err := parser.commandType()
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(aCommand{}, lCommand{}, cCommand{})); diff != "" {
				t.Errorf(diff)
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
			parser := p()
			got, err := parser.toACommand(tc.raw)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(aCommand{})); diff != "" {
				t.Errorf("want = %#v, got = %#v, \ndiff: %s", tc.want, got, diff)
			}
		})
	}
}

func TestToLCommand(t *testing.T) {
	testCases := map[string]struct {
		raw  string
		want *lCommand
	}{
		"ok": {"(hoge)", l(lSymbol("hoge"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			parser := p()
			got, err := parser.toLCommand(tc.raw)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(lCommand{})); diff != "" {
				t.Errorf("want = %#v, but got = %#v\ndiff=%v", tc.want, got, diff)
			}
		})
	}
}

func TestToCCommand(t *testing.T) {
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

			parser := p()
			got, err := parser.toCCommand(tc.in)
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(cCommand{})); diff != "" {
				t.Errorf("want != got\ndiff=%s", diff)
			}
		})
	}
}

type Option[T any] func(val T)
type ParserOption Option[*Parser]
type ACommandOption func(val *aCommand)
type LCommandOption Option[*lCommand]

func l(opts ...LCommandOption) *lCommand {
	lCommand := &lCommand{
		symbol: "",
	}

	for _, opt := range opts {
		opt(lCommand)
	}

	return lCommand
}

func lSymbol(v string) LCommandOption {
	return func(val *lCommand) {
		val.symbol = v
	}
}

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

func p(opts ...ParserOption) Parser {
	parser := &Parser{
		commands:    []string{},
		currentIdx:  0,
		symbolTable: symtable.New(),
	}

	for _, opt := range opts {
		opt(parser)
	}

	return *parser
}

func commands(vals ...string) ParserOption {
	return func(val *Parser) {
		val.commands = vals
	}
}

func currentIdx(idx int) ParserOption {
	return func(val *Parser) {
		val.currentIdx = idx
	}
}

func s[T any](vals ...T) []T {
	return vals
}
