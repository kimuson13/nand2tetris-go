package parser

import (
	"assembler/code"
	"strconv"
)

type ACommand struct {
	address int
	symbol  string
}

func (a *ACommand) parse() (code.Command, error) {
	return &code.ACommand{
		Address: a.address,
		Symbol:  a.symbol,
	}, nil
}

func isACommand(raw string) bool {
	head := raw[0]
	return head == '@'
}

func toACommand(raw string) (*ACommand, error) {
	val := string(raw[1:])
	i, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}

	return &ACommand{address: i, symbol: ""}, nil
}
