package process

import (
	"assembler/parser"
	"errors"
	"fmt"
)

var (
	ErrInvalidArgsLength = errors.New("invalid args length. this needs a file path only")
)

func Run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("process error: %w", ErrInvalidArgsLength)
	}

	_, err := parser.Parse(args[0])
	if err != nil {
		return fmt.Errorf("process error: %w", err)
	}
	return nil
}
