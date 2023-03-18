package codewriter

type Segment int

const (
	CONSTANT Segment = iota
	ARGUMENT
	LOCAL
	THAT
	THIS
	POINTER
	TEMP
	STATIC
)
