package codewriter

type Command interface {
	convert() ([]byte, error)
}

type ArithmeticKind int

const (
	ADD ArithmeticKind = iota
)

type Arithmetic struct {
	Kind ArithmeticKind
}

func (a Arithmetic) convert() ([]byte, error) {
	return nil, nil
}
