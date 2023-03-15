package parser_test

import (
	"testing"
	"vmtranslator/codewriter"
	"vmtranslator/parser"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	in := "// comment \n\npush constant 6\nadd // comment"

	if _, err := parser.New(in); err != nil {
		t.Error(err)
	}
}

func TestParse(t *testing.T) {
	in := "push constant 6\nadd"
	want := []codewriter.Command{
		codewriter.Push{Segment: codewriter.CONSTANT, Index: 6},
		codewriter.Arithmetic{Kind: codewriter.ADD},
	}

	p, err := parser.New(in)
	if err != nil {
		t.Fatal(err)
	}

	got, err := p.Parse()
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
}
