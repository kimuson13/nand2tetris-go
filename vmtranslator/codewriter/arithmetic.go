package codewriter

import (
	"fmt"
	"math/rand"
	"strconv"
)

type ArithmeticKind int

const (
	ADD ArithmeticKind = iota
	SUB
	NEGATIVE
	EQUAL
	GREATER_THAN
	LOWER_THAN
	AND
	OR
	NOT
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
	case SUB:
		return a.genSub()
	case NEGATIVE:
		return a.genNegative()
	case EQUAL, GREATER_THAN, LOWER_THAN:
		return a.genCompare()
	case AND:
		return a.genAnd()
	case OR:
		return a.genOr()
	case NOT:
		return a.genNot()
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

func (a Arithmetic) genSub() []byte {
	const subAsm = `
@SP
A=M
A=A-1
D=M
A=A-1
M=M-D
@SP
M=M-1
`
	return []byte(subAsm)
}

func (a Arithmetic) genNegative() []byte {
	const negativeAsm = `
@SP
A=M
A=A-1
M=-M
`
	return []byte(negativeAsm)
}

func (a Arithmetic) genCompare() []byte {
	const compareAsm = `
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@TRUE%[1]s
%[2]s
@SP
A=M
M=0
@NEXT%[1]s
0;JMP
(TRUE%[1]s)
@SP
A=M
M=0
M=-1
(NEXT%[1]s)
@SP
M=M+1
`
	compareAddressId := strconv.Itoa(rand.Intn(1000000))
	var compareStmt string
	switch a.Kind {
	case EQUAL:
		compareStmt = "D;JEQ"
	case GREATER_THAN:
		compareStmt = "D;JGT"
	case LOWER_THAN:
		compareStmt = "D;JLT"
	}
	return []byte(fmt.Sprintf(compareAsm, compareAddressId, compareStmt))
}

func (a Arithmetic) genAnd() []byte {
	const andAsm = `
@SP
A=A-1
D=M
A=A-1
M=M&D
@SP
M=M-1
`
	return []byte(andAsm)
}

func (a Arithmetic) genOr() []byte {
	const orAsm = `
@SP
A=A-1
D=M
A=A-1
M=M|D
@SP
M=M-1
`
	return []byte(orAsm)
}

func (a Arithmetic) genNot() []byte {
	const notAsm = `
@SP
A=M
A=A-1
M=!M
`
	return []byte(notAsm)
}
