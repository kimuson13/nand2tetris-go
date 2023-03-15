package parser

type commandType string

const (
	ADD  commandType = "add"
	PUSH commandType = "push"
)

type command int

const (
	INVALID command = iota
	C_ARITHMETIC
	C_PUSH
)
