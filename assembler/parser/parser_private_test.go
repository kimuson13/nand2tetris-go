package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

type Option[T any] func(val T)
type ParserOption Option[*Parser]

func p(opts ...ParserOption) Parser {
	parser := &Parser{
		commands:   []string{},
		currentIdx: 0,
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
