package codewriter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPop_convert(t *testing.T) {
	testCases := map[string]struct {
		pop       Pop
		expected  string
		expectErr bool
	}{
		"Static": {
			pop:       Pop{Segment: STATIC, Index: 3, FileName: "Test"},
			expected:  "\n@SP\nA=M\nA=A-1\nD=M\n@TEST_3\nM=D\n@SP\nM=M-1\n",
			expectErr: false,
		},
		"Local": {
			pop:       Pop{Segment: LOCAL, Index: 2},
			expected:  "\n@2\nD=A\n@LCL\nA=M\nD=D+A\n@temp\nM=D\n@SP\nA=M\nA=A-1\nD=M\n@temp\nA=M\nM=D\n@SP\nM=M-1\n",
			expectErr: false,
		},
		"Invalid": {
			pop:       Pop{Segment: CONSTANT, Index: 2},
			expected:  "",
			expectErr: true,
		},
	}

	for name, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel() // Mark this test as parallel
			result, err := tc.pop.convert()
			if diff := cmp.Diff(string(result), tc.expected); diff != "" {
				t.Errorf("convert() mismatch (-want +got):\n%s", diff)
			}
			if tc.expectErr && err == nil {
				t.Errorf("convert() expected error but got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("convert() unexpected error: %v", err)
			}
		})
	}
}

func TestPop_genStatic(t *testing.T) {
	pop := Pop{Segment: STATIC, Index: 2, FileName: "Test"}
	expected := "\n@SP\nA=M\nA=A-1\nD=M\n@TEST_2\nM=D\n@SP\nM=M-1\n"

	result := pop.genStatic()
	if diff := cmp.Diff(string(result), expected); diff != "" {
		t.Errorf("genStatic() mismatch (-want +got):\n%s", diff)
	}
}

func TestPop_genMemoryAccess(t *testing.T) {
	testCases := map[string]struct {
		pop      Pop
		expected string
	}{
		"Local":    {Pop{Segment: LOCAL, Index: 2}, "\n@2\nD=A\n@LCL\nA=M\nD=D+A\n@temp\nM=D\n@SP\nA=M\nA=A-1\nD=M\n@temp\nA=M\nM=D\n@SP\nM=M-1\n"},
		"Argument": {Pop{Segment: ARGUMENT, Index: 2}, "\n@2\nD=A\n@ARG\nA=M\nD=D+A\n@temp\nM=D\n@SP\nA=M\nA=A-1\nD=M\n@temp\nA=M\nM=D\n@SP\nM=M-1\n"},
		"That":     {Pop{Segment: THAT, Index: 2}, "\n@2\nD=A\n@THAT\nA=M\nD=D+A\n@temp\nM=D\n@SP\nA=M\nA=A-1\nD=M\n@temp\nA=M\nM=D\n@SP\nM=M-1\n"},
		"This":     {Pop{Segment: THIS, Index: 2}, "\n@2\nD=A\n@THIS\nA=M\nD=D+A\n@temp\nM=D\n@SP\nA=M\nA=A-1\nD=M\n@temp\nA=M\nM=D\n@SP\nM=M-1\n"},
		"Temp":     {Pop{Segment: TEMP, Index: 2}, "\n@2\nD=A\n@5\nD=D+A\n@temp\nM=D\n@SP\nA=M\nA=A-1\nD=M\n@temp\nA=M\nM=D\n@SP\nM=M-1\n"},
		"Pointer":  {Pop{Segment: POINTER, Index: 2}, "\n@2\nD=A\n@3\nD=D+A\n@temp\nM=D\n@SP\nA=M\nA=A-1\nD=M\n@temp\nA=M\nM=D\n@SP\nM=M-1\n"},
	}

	for name, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel() // Mark this test as parallel
			result := tc.pop.genMemoryAccess()
			if diff := cmp.Diff(string(result), tc.expected); diff != "" {
				t.Errorf("genMemoryAccess() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
