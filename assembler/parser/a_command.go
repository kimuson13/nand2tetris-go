package parser

import (
	"assembler/code"
	"strconv"
)

type ACommand struct {
	value  int
	symbol string
}

func (a *ACommand) parse() code.Command {
	return nil
}

func toACommand(raw string) (*ACommand, error) {
	val := string(raw[0:])
	i, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}

	return &ACommand{value: i, symbol: val}, nil
}
