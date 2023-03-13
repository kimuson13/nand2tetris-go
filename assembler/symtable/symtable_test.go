package symtable_test

import (
	"assembler/symtable"
	"testing"
)

func TestSymTable_AddEntry(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		symbol  string
		address int
		wantErr bool
	}{
		"ok": {"hoge", 123, noErr},
		"ng": {symtable.ARG, 122, wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			symTable := symtable.New()
			err := symTable.AddEntry(tc.symbol, tc.address)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("no error: sym = %s, address = %d", tc.symbol, tc.address)
			}
		})
	}
}

func TestSymTable_GetAddress(t *testing.T) {
	const wantErr, noErr = true, false
	testCases := map[string]struct {
		symbol  string
		want    int
		wantErr bool
	}{
		"ok": {symtable.SCREEN, 16384, noErr},
		"ng": {"hoge", 0, wantErr},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			symtable := symtable.New()
			got, err := symtable.GetAddress(tc.symbol)
			if err != nil && !tc.wantErr {
				t.Error(err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("no error: sym: %s", tc.symbol)
			}

			if got != tc.want {
				t.Errorf("want = %d, got = %d", tc.want, got)
			}
		})
	}
}

func TestSymTable_Contains(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want bool
	}{}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			symTable := symtable.New()
			if got := symTable.Contains(tc.in); got != tc.want {
				t.Errorf("want = %v, but got = %v", tc.want, got)
			}
		})
	}
}
