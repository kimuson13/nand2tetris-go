package parser_test

import (
	"testing"
	"vmtranslator/parser"
)

func TestNew(t *testing.T) {
	in := "// comment \n\npush constant 6\nadd // comment"

	if _, err := parser.New(in); err != nil {
		t.Error(err)
	}
}
