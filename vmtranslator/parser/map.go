package parser

import (
	"errors"
	"fmt"
	"vmtranslator/codewriter"
)

var (
	ErrNoSuchKey = errors.New("no such a key")
)

func mapCommandTypeToArithmeticKind(cType commandType) (codewriter.ArithmeticKind, error) {
	mp := map[commandType]codewriter.ArithmeticKind{
		ADD: codewriter.ADD,
	}

	val, ok := mp[cType]
	if !ok {
		return 0, fmt.Errorf("mapCommandTypeToCodeWriterKind: %w", ErrNoSuchKey)
	}

	return val, nil
}

func mapSegment(raw string) (codewriter.Segment, error) {
	mp := map[string]codewriter.Segment{
		"constant": codewriter.CONSTANT,
	}

	val, ok := mp[raw]
	if !ok {
		return 0, fmt.Errorf("mapSegment error: %w", ErrNoSuchKey)
	}

	return val, nil
}
