package codewriter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPush_genMemmoryAccess(t *testing.T) {
	testCases := map[string]struct {
		push     Push
		expected string
	}{
		"local":    {Push{Segment: LOCAL, Index: 1}, "\n@1\nD=A\n@LCL\nA=M\nA=A+D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		"argument": {Push{Segment: ARGUMENT, Index: 2}, "\n@2\nD=A\n@ARG\nA=M\nA=A+D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		"push":     {Push{Segment: THIS, Index: 3}, "\n@3\nD=A\n@THIS\nA=M\nA=A+D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		"that":     {Push{Segment: THAT, Index: 4}, "\n@4\nD=A\n@THAT\nA=M\nA=A+D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		"temp":     {Push{Segment: TEMP, Index: 5}, "\n@5\nD=A\n@5\nA=A+D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		"pointer":  {Push{Segment: POINTER, Index: 6}, "\n@6\nD=A\n@3\nA=A+D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
	}

	for name, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel() // Mark this test as parallel
			result := tc.push.genMemoryAccess()
			if diff := cmp.Diff(string(result), tc.expected); diff != "" {
				t.Errorf("genMemoryAccess() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPush_genAsm(t *testing.T) {
	testCases := map[string]struct {
		push     Push
		expected string
	}{
		"Constant": {Push{Segment: CONSTANT, Index: 7}, "\n@7\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		"Argument": {Push{Segment: ARGUMENT, Index: 2}, "\n@2\nD=A\n@ARG\nA=M\nA=A+D\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		"Static":   {Push{Segment: STATIC, Index: 3, FileName: "Test"}, "\n@TEST_3\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
	}

	for name, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel() // Mark this test as parallel
			result := tc.push.genAsm()
			if diff := cmp.Diff(string(result), tc.expected); diff != "" {
				t.Errorf("genAsm() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPush_genConstant(t *testing.T) {
	push := Push{Segment: CONSTANT, Index: 10}
	expected := "\n@10\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"

	result := push.genConstant()
	if diff := cmp.Diff(string(result), expected); diff != "" {
		t.Errorf("genConstant() mismatch (-want +got):\n%s", diff)
	}
}

func TestPush_genStatic(t *testing.T) {
	push := Push{Segment: STATIC, Index: 1, FileName: "Test"}
	expected := "\n@TEST_1\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"

	result := push.genStatic()
	if diff := cmp.Diff(string(result), expected); diff != "" {
		t.Errorf("genStatic() mismatch (-want +got):\n%s", diff)
	}
}
