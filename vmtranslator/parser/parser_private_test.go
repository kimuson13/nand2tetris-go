package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	in := "// comment \n\npush constant 6\nadd // comment"
	want := Parser{
		commands:       s("push constant 6", "add"),
		currentCommand: s("push", "constant", "6"),
	}

	got, err := New(in)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Parser{})); diff != "" {
		t.Errorf("in: %s\n%s", in, diff)
	}
}

func TestGetCommands(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    []string
		wantErr bool
	}{
		"get_commands":   {"// comment\n\nadd // comment", s("add"), noErr},
		"invalid_syntax": {"add // comment // comment", []string{}, wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := getCommands(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("no err: %v", tc.in)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error(diff)
			}
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
		"command":              {"add", "add", noErr},
		"comment_line":         {"// comment", "", noErr},
		"empty_line":           {"", "", noErr},
		"command_with_comment": {"push constant 6 // comment", "push constant 6", noErr},
		"too_many_comment_lit": {"add // comment // comment", "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := getCommand(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("no err: %v", tc.in)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestIsNotCommand(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want bool
	}{
		"empty_line":   {"", true},
		"comment_line": {"// comment", true},
		"command_line": {"push constant 6", false},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
		})

		if got := isNotCommand(tc.in); got != tc.want {
			t.Errorf("want = %v, but got =%v", tc.want, got)
		}
	}
}

func TestTrimComment(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		in      string
		want    string
		wantErr bool
	}{
		"no_comment":           {"year", "year", noErr},
		"{val}_space_comment":  {"year // comment", "year", noErr},
		"{val}_comment":        {"year//comment", "year", noErr},
		"too_many_comment_lit": {"year //comment // comment", "", wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := trimComment(tc.in)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("no err: %v", tc.in)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff:\n%s", diff)
			}
		})
	}
}

func TestParser_hasMoreCommand(t *testing.T) {
	testCases := map[string]struct {
		opts []ParserOption
		want bool
	}{
		"hasMore":     {s(currentIdx(0), commands("hoge", "huga")), true},
		"not_hasMore": {s(currentIdx(5), commands("hoge", "huga")), false},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			parser := genParser(tc.opts...)
			if got := parser.hasMoreCommand(); got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}

func s[T any](val ...T) []T {
	return val
}

type ParserOption func(val *Parser)

func genParser(opts ...ParserOption) Parser {
	parser := Parser{}

	for _, opt := range opts {
		opt(&parser)
	}

	return parser
}

func commands(v ...string) ParserOption {
	return func(val *Parser) {
		slice := s(v...)
		val.commands = slice
	}
}

func currentIdx(v int) ParserOption {
	return func(val *Parser) {
		val.currentIdx = v
	}
}

// test format
/*
const wantErr, noErr = true, false
testCases := map[string]struct {
	}{}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
		})
	}
*/
/* test case with wantErr
if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("no err: %v", )
			}
*/
