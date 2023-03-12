package parser

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrepare(t *testing.T) {
	testCases := map[string]struct {
		fileVal string
		want    Parser
	}{
		"one_line":        {"test", p(commands("test"))},
		"many_lines":      {"test2\r\nhoge", p(commands("test2", "hoge"))},
		"include_space":   {"test     \r\nhuga", p(commands("test", "huga"))},
		"include_comment": {"//comment\r\n@123", p(commands("@123"))},
		"inline_comment":  {"@123 //123\r\n(hoge)", p(commands("@123", "(hoge)"))},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f, err := os.CreateTemp("./", "")
			if err != nil {
				t.Fatal(err)
			}
			f.Write([]byte(tc.fileVal))

			got, err := prepare(f.Name())
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(Parser{})); diff != "" {
				t.Errorf("want = %#v, got = %#v, \ndiff: %s", tc.want, got, diff)
			}

			os.Remove(f.Name())
		})
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
		"a_command": {s("@123"), a(aValue(123))},
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

			if diff := cmp.Diff(tc.want, got, cmp.AllowUnexported(ACommand{}, LCommand{}, CCommand{})); diff != "" {
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
