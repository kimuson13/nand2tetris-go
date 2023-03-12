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
		asmPath string
		want    string
	}{
		"MaxL":  {"../tests/asm/MaxL.asm", getBody(t, "../tests/hack/MaxL.hack")},
		"PongL": {"../tests/asm/PongL.asm", getBody(t, "../tests/hack/PongL.hack")},
		"RectL": {"../tests/asm/RectL.asm", getBody(t, "../tests/hack/RectL.hack")},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dir, fileName := filepath.Split(tc.asmPath)
			ext := filepath.Ext(tc.asmPath)
			genHackPath := filepath.Join(dir, strings.ReplaceAll(fileName, ext, ".hack"))

			if err := process.Run([]string{tc.asmPath}); err != nil {
				t.Fatal(err)
			}

			got := getBody(t, genHackPath)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Error(diff)
			}

			if err := os.Remove(genHackPath); err != nil {
				t.Error(err)
			}
		})
	}
}

func getBody(t *testing.T, path string) string {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}
