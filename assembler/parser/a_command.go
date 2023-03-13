package parser

import (
	"assembler/code"
)

type aCommand struct {
	address int
	symbol  string
}

func (a *aCommand) parse() (code.Command, error) {
	return &code.ACommand{
		Address: a.address,
		Symbol:  a.symbol,
	}, nil
}

func isACommand(raw string) bool {
	head := raw[0]
	return head == '@'
}
