package codewriter

import (
	"fmt"
	"strings"
)

type Pop struct {
	FileName string
	Segment  Segment
	Index    int
}

func (p Pop) convert() ([]byte, error) {
	switch p.Segment {
	case STATIC:
		return p.genStatic(), nil
	case LOCAL, ARGUMENT, THAT, THIS, TEMP, POINTER:
		return p.genMemoryAccess(), nil
	}

	return nil, fmt.Errorf("pop convert error: %w", ErrCanNotConvert)
}

func (p Pop) genStatic() []byte {
	const asm = `
@SP
A=M
A=A-1
D=M
@%s_%d
M=D
@SP
M=M-1
`
	upperFileName := strings.ToUpper(p.FileName)
	return []byte(fmt.Sprintf(asm, upperFileName, p.Index))
}

func (p Pop) genMemoryAccess() []byte {
	const asm = `
@%d
D=A
%s
D=D+A
@temp
M=D
@SP
A=M
A=A-1
D=M
@temp
A=M
M=D
@SP
M=M-1
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
		line = "@5"
	case POINTER:
		line = "@3"
	}

	return []byte(fmt.Sprintf(asm, p.Index, line))
}
