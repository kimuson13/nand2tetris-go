package parser

import "assembler/code"

type LCommand struct {
	symbol string
}

func (l *LCommand) parse() (code.Command, error) {
	return &code.LCommand{
		Symbol: l.symbol,
	}, nil
}

func isLCommand(raw string) bool {
	head := raw[0]
	tail := raw[len(raw)-1]
	return head == '(' && tail == ')'
}

func toLCommand(raw string) (*LCommand, error) {
	val := string(raw[1 : len(raw)-1])

	return &LCommand{symbol: val}, nil
}
