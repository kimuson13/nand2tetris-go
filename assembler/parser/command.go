package parser

import (
	"assembler/code"
)

type Command interface {
	parse() (code.Command, error)
}
