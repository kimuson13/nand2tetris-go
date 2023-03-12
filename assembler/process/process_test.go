package process_test

import (
	"assembler/process"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()
	in := []string{"test.asm"}
	want := "0000000001100100\r\n1111110000010111\r\n"

	if err := process.Run(in); err != nil {
		t.Error(err)
	}

	b, err := os.ReadFile("./test.hack")
	if err != nil {
		t.Error(err)
	}

	if got := string(b); got != want {
		t.Errorf("want = %v, but got = %v", want, got)
	}
	if err := os.Remove("./test.hack"); err != nil {
		t.Error(err)
	}
}

func setUp(t *testing.T) func() {
	t.Helper()

	f, err := os.Create("test.asm")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	s := "//comment\r\n@100\r\n(HOGE)\r\nD=M;JMP //comment2"
	if _, err := f.Write([]byte(s)); err != nil {
		t.Fatal(err)
	}

	return func() {
		if err := os.Remove(f.Name()); err != nil {
			t.Error(err)
		}
	}
}
