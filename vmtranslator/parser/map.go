package parser

import (
	"fmt"
	"vmtranslator/codewriter"
)

func mapCommandTypeToArithmeticKind(cType commandType) (codewriter.ArithmeticKind, error) {
	mp := map[commandType]codewriter.ArithmeticKind{
		ADD: codewriter.ADD,
	}

	val, ok := mp[cType]
	if !ok {
		return 0, fmt.Errorf("mapCommandTypeToCodeWriterKind: %w", ErrInvalidCommand)
	}

	return val, nil
}
