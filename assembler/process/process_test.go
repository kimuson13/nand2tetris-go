package process_test

import (
	"assembler/process"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	testCases := map[string]struct {
		asmPath  string
		wantPath string
	}{
		"MaxL":  {"MaxL.asm", "MaxL.hack"},
		"PongL": {"PongL.asm", "PongL.hack"},
		"RectL": {"RectL.asm", "RectL.hack"},
		"Add":   {"Add.asm", "Add.hack"},
		"Max":   {"Max.asm", "Max.hack"},
		"Pong":  {"Pong.asm", "Pong.hack"},
		"Rect":  {"Rect.asm", "Rect.hack"},
		"Prog":  {"Prog.asm", "Prog.hack"},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			asmPath := joinAsm(tc.asmPath)

			dir, fileName := filepath.Split(asmPath)
			ext := filepath.Ext(tc.asmPath)
			genHackPath := filepath.Join(dir, strings.ReplaceAll(fileName, ext, ".hack"))

			if err := process.Run([]string{asmPath}); err != nil {
				t.Fatal(err)
			}

			got := getBody(t, genHackPath)
			want := getBody(t, joinHack(tc.wantPath))
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("want:\n%s\ngot:\n%s\ndiff:\n%s", want, got, diff)
			}

			if err := os.Remove(genHackPath); err != nil {
				t.Error(err)
			}
		})
	}
}

func joinAsm(path string) string {
	return filepath.Join("../tests/asm/", path)
}

func joinHack(path string) string {
	return filepath.Join("../tests/hack/", path)
}

func getBody(t *testing.T, path string) string {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}
