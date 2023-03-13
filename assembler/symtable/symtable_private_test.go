package symtable

import "testing"

func TestGenDefinedSymbol(t *testing.T) {
	wantLength := 23
	got := genDefinedSymbol()

	if gotLength := len(got); gotLength != wantLength {
		t.Errorf("length not much: want = %d, but got = %d", wantLength, gotLength)
	}
}
