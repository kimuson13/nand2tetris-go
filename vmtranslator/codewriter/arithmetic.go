package codewriter

import "fmt"

type ArithmeticKind int

const (
	ADD ArithmeticKind = iota
)

type Arithmetic struct {
	Kind ArithmeticKind
}

func (a Arithmetic) convert() ([]byte, error) {
	b := a.genAsm()
	if b == nil {
		return nil, fmt.Errorf("arithmetic error: %w", ErrCanNotConvert)
	}
	return b, nil
}

func (a Arithmetic) genAsm() []byte {
	switch a.Kind {
	case ADD:
		return a.genAdd()
	}

	return nil
}

func (a Arithmetic) genAdd() []byte {
	const addAsm = `
@SP
A=M
A=A-1
D=M
A=A-1
M=M+D
@SP
M=M-1
`

	return []byte(addAsm)
}
