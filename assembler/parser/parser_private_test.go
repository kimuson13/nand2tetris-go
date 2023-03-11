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
		"one_line":      {"test", p(commands("test"))},
		"many_lines":    {"test2\r\nhoge", p(commands("test2", "hoge"))},
		"include_space": {"test     \r\nhuga", p(commands("test", "huga"))},
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

type Option[T any] func(val T)

func p(opts ...Option[*Parser]) Parser {
	parser := &Parser{
		commands:   []string{},
		currentIdx: 0,
	}

	for _, opt := range opts {
		opt(parser)
	}

	return *parser
}

func commands(vals ...string) Option[*Parser] {
	return func(val *Parser) {
		val.commands = vals
	}
}

func currentIdx(idx int) Option[*Parser] {
	return func(val *Parser) {
		val.currentIdx = idx
	}
}

func s[T any](vals ...T) []T {
	return vals
}
