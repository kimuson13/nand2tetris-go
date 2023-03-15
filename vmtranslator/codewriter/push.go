package codewriter

import (
	"errors"
	"fmt"
)

type Push struct {
	Segment Segment
	Index   int
}

var (
	ErrCanNotConvert = errors.New("can not convert")
)

func (p Push) convert() ([]byte, error) {
	b := p.genAsm()
	if b == nil {
		return nil, fmt.Errorf("push error: %w, segment = %v, index = %d", ErrCanNotConvert, p.Segment, p.Index)
	}

	return b, nil
}

func (p Push) genAsm() []byte {
	switch p.Segment {
	case CONSTANT:
		return p.genConstant()
	}

	return nil
}

func (p Push) genConstant() []byte {
	const constantAsm = `
@%d
D=A
@SP
A=M
M=D
@SP
M=M+1
`

	asm := fmt.Sprintf(constantAsm, p.Index)
	return []byte(asm)
}
