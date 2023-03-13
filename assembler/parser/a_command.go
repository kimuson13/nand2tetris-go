package parser

import (
	"assembler/code"
	"strconv"
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

func toACommand(raw string) (*aCommand, error) {
	val := string(raw[1:])
	i, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}

	return &aCommand{address: i, symbol: ""}, nil
}
