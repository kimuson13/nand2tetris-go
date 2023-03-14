package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
