package parser_test

import (
	"assembler/code"
	"assembler/parser"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser_SymbolicLink(t *testing.T) {
	f := setUp(t)

	p, err := parser.New(f.Name())
	if err != nil {
		t.Error(err)
	}
	if err := p.SymbolicLink(); err != nil {
		t.Error(err)
	}

	os.Remove(f.Name())
}

func TestParser_Parse(t *testing.T) {
	want := []code.Command{
		&code.ACommand{Address: 123},
		&code.LCommand{Symbol: "HOGE"},
		&code.CCommand{
			Dest: ptr(code.DEST_M),
			Comp: code.COMP_M_ADD_1,
			Jump: ptr(code.JUMP)},
		&code.CCommand{
			Dest: ptr(code.DEST_D),
			Comp: code.COMP_D_MINUS_M,
		},
	}
	f := setUp(t)

	p, err := parser.New(f.Name())
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
	os.Remove(f.Name())
}

func setUp(t *testing.T) *os.File {
	f, err := os.CreateTemp("./", "")
	if err != nil {
		t.Fatal(err)
	}
	in := "//hogehoge\r\n@123 //hoge\r\n(HOGE)\r\nM=M+1;JMP\r\nD=D-M"
	if _, err := f.Write([]byte(in)); err != nil {
		t.Fatal(err)
	}

	return f
}

func s[T any](val ...T) []T {
	return val
}

func ptr[T any](val T) *T {
	return &val
}
