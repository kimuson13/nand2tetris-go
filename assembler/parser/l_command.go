package parser

import "assembler/code"

type LCommand struct {
	value  int
	symbol string
}

func (l *LCommand) parse() code.Command {
	return nil
}

func isLCommand(raw string) bool {
	head := raw[0]
	return head == '('
}

func toLCommand(raw string) (*LCommand, error) {
	val := string(raw[0 : len(raw)-1])

	return &LCommand{value: 0, symbol: val}, nil
}
