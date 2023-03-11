package parser

import "assembler/code"

type Command interface {
	parse() code.Command
}

type CCommand struct {
	dest string
	comp string
	jump string
}

func (c *CCommand) parse() code.Command {
	return nil
}

type ACommand struct {
	value  int
	symbol string
}

func (a *ACommand) parse() code.Command {

	return nil
}

type LCommand struct {
	value  int
	symbol string
}

func (l *LCommand) parse() code.Command {
	return nil
}
