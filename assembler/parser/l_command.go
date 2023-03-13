package parser

import "assembler/code"

type lCommand struct {
	symbol string
}

func (l *lCommand) parse() (code.Command, error) {
	return &code.LCommand{
		Symbol: l.symbol,
	}, nil
}

func isLCommand(raw string) bool {
	head := raw[0]
	tail := raw[len(raw)-1]
	return head == '(' && tail == ')'
}

func toLCommand(raw string) (*lCommand, error) {
	val := string(raw[1 : len(raw)-1])

	return &lCommand{symbol: val}, nil
}
