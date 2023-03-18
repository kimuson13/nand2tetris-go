package codewriter

import (
	"errors"
	"fmt"
)

type Push struct {
	FileName string
	Segment  Segment
	Index    int
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
	case ARGUMENT, LOCAL, THAT, THIS, POINTER, TEMP:
		return p.genMemoryAccess()
	case STATIC:
		return p.genStatic()
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

func (p Push) genMemoryAccess() []byte {
	const memoryAccess = `
@%d
D=A
%s
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1
`

	var line string
	switch p.Segment {
	case LOCAL:
		line = "@LCL\nA=M"
	case ARGUMENT:
		line = "@ARG\nA=M"
	case THAT:
		line = "@THAT\nA=M"
	case THIS:
		line = "@THIS\nA=M"
	case TEMP:
		line = "@5\n"
	case POINTER:
		line = "@3\n"
	}

	return []byte(fmt.Sprintf(memoryAccess, p.Index, line))
}

func (p Push) genStatic() []byte {
	const asm = `
@%s_%d
D=M
@SP
A=M
M=D
@SP
M=M+1
`
	return []byte(fmt.Sprintf(asm, p.FileName, p.Index))
}
