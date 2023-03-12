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

	commands, err := parser.Parse(args[0])
	if err != nil {
		return fmt.Errorf("process error: %w", err)
	}

	binaryLines := make([]string, 0, len(commands))
	for _, command := range commands {
		bLine, err := command.Convert()
		if err != nil {
			return fmt.Errorf("process error: %w", err)
		}
		if bLine != "" {
			binaryLines = append(binaryLines, bLine)
		}
	}
	return nil
}
