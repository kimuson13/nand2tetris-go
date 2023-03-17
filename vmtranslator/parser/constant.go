package parser

type commandType string

const (
	ADD          commandType = "add"
	SUB          commandType = "sub"
	NEGATIVE     commandType = "neg"
	EQUAL        commandType = "eq"
	GREATER_THAN commandType = "gt"
	LOWER_THAN   commandType = "lt"
	AND          commandType = "and"
	OR           commandType = "or"
	NOT          commandType = "not"
	PUSH         commandType = "push"
	POP          commandType = "pop"
)

type command int

const (
	INVALID command = iota
	C_ARITHMETIC
	C_PUSH
	C_POP
)
