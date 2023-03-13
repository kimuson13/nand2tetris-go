package parser

import (
	"assembler/code"
	"regexp"
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

var (
	aCommandDirectAddressRegExp = regexp.MustCompile(`^@[0-9]+$`)
	aCommandSymbolRegExp        = regexp.MustCompile(`^@[a-zA-Z\.\$_:]{1}[0-9a-zA-Z\.\$_:]*$`)
)

func isACommand(raw string) bool {
	return aCommandDirectAddressRegExp.MatchString(raw) || aCommandSymbolRegExp.MatchString(raw)
}
