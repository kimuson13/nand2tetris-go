package file_test

import (
	"assembler/file"
	"os"
	"testing"
)

func TestCreateHack(t *testing.T) {
	path := "./hoge.asm"
	lines := []string{"hoge", "huga"}
	want := "hoge\nhuga\n"

	if err := file.CreateHack(path, lines); err != nil {
		t.Fatal(err)
	}

	b, err := os.ReadFile("./hoge.hack")
	if err != nil {
		t.Error(err)
	}

	if got := string(b); got != want {
		t.Errorf("want = %s, got = %s", want, got)
	}
	os.Remove("./hoge.hack")
}
