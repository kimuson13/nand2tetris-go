package parser

import "vmtranslator/parser/codewriter"

type command interface {
	parse() (codewriter.Command, error)
}

type arithmetic struct {
}

func (a arithmetic) parse() (codewriter.Command, error) {
	return nil, nil
}

type push struct{}

func (p push) parse() (codewriter.Command, error) {
	return nil, nil
}
