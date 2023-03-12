package parser_test

import (
	"assembler/code"
	"assembler/parser"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	f, err := os.CreateTemp("./", "")
	if err != nil {
		t.Fatal(err)
	}
	in := "//hogehoge\r\n@123 //hoge\r\n(HOGE)\r\nM=M+1;JMP"
	want := []code.Command{
		&code.ACommand{Value: 123},
		&code.LCommand{Symbol: "HOGE"},
		&code.CCommand{
			Dest: ptr(code.DEST_M),
			Comp: code.COMP_M_ADD_1,
			Jump: ptr(code.JUMP)},
	}
	f.Write([]byte(in))
	got, err := parser.Parse(f.Name())
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Error(diff)
	}
	os.Remove(f.Name())
}

func s[T any](val ...T) []T {
	return val
}

func ptr[T any](val T) *T {
	return &val
}
