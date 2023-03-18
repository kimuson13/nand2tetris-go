package codewriter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArithmetic_convert(t *testing.T) {
	testCases := map[string]struct {
		arithmetic Arithmetic
		expected   string
	}{
		"Add": {
			arithmetic: Arithmetic{Kind: ADD},
			expected:   "\n@SP\nA=M\nA=A-1\nD=M\nA=A-1\nM=M+D\n@SP\nM=M-1\n",
		},
		"Sub": {
			arithmetic: Arithmetic{Kind: SUB},
			expected:   "\n@SP\nA=M\nA=A-1\nD=M\nA=A-1\nM=M-D\n@SP\nM=M-1\n",
		},
		"Negative": {
			arithmetic: Arithmetic{Kind: NEGATIVE},
			expected:   "\n@SP\nA=M\nA=A-1\nM=-M\n",
		},
		"And": {
			arithmetic: Arithmetic{Kind: AND},
			expected:   "\n@SP\nA=M\nA=A-1\nD=M\nA=A-1\nM=M&D\n@SP\nM=M-1\n",
		},
		"Or": {
			arithmetic: Arithmetic{Kind: OR},
			expected:   "\n@SP\nA=M\nA=A-1\nD=M\nA=A-1\nM=M|D\n@SP\nM=M-1\n",
		},
		"Not": {
			arithmetic: Arithmetic{Kind: NOT},
			expected:   "\n@SP\nA=M\nA=A-1\nM=!M\n",
		},
	}

	for name, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel() // Mark this test as parallel
			result, err := tc.arithmetic.convert()
			if err != nil {
				t.Errorf("convert() unexpected error: %v", err)
			}
			if diff := cmp.Diff(string(result), tc.expected); diff != "" {
				t.Errorf("convert() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestArithmetic_genCompare(t *testing.T) {
	testCases := map[string]struct {
		arithmetic Arithmetic
	}{
		"Equal": {
			arithmetic: Arithmetic{Kind: EQUAL},
		},
		"Greater_Than": {
			arithmetic: Arithmetic{Kind: GREATER_THAN},
		},
		"Lower_Than": {
			arithmetic: Arithmetic{Kind: LOWER_THAN},
		},
	}

	for name, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel() // Mark this test as parallel
			result := tc.arithmetic.genCompare()
			if len(result) == 0 {
				t.Errorf("genCompare() returned an empty byte slice")
			}
		})
	}
}
