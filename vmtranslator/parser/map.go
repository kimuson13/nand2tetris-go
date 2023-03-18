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
		ADD:          codewriter.ADD,
		SUB:          codewriter.SUB,
		NEGATIVE:     codewriter.NEGATIVE,
		EQUAL:        codewriter.EQUAL,
		GREATER_THAN: codewriter.GREATER_THAN,
		LOWER_THAN:   codewriter.LOWER_THAN,
		AND:          codewriter.AND,
		OR:           codewriter.OR,
		NOT:          codewriter.NOT,
	}

	val, ok := mp[cType]
	if !ok {
		return 0, fmt.Errorf("mapCommandTypeToCodeWriterKind: %w", ErrNoSuchKey)
	}

	return val, nil
}

func mapSegment(raw string) (codewriter.Segment, error) {
	mp := map[string]codewriter.Segment{
		"argument": codewriter.ARGUMENT,
		"local":    codewriter.LOCAL,
		"static":   codewriter.STATIC,
		"constant": codewriter.CONSTANT,
		"this":     codewriter.THIS,
		"that":     codewriter.THAT,
		"pointer":  codewriter.POINTER,
		"temp":     codewriter.TEMP,
	}

	val, ok := mp[raw]
	if !ok {
		return 0, fmt.Errorf("mapSegment error: %w", ErrNoSuchKey)
	}

	return val, nil
}
