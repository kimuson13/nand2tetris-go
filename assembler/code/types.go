package code

import "errors"

var (
	ErrNoSuchDest = errors.New("no such dest binary")
	ErrNoSuchComp = errors.New("no such comp binary")
	ErrNoSuchJump = errors.New("no such jump binary")
)

type Dest string
type Comp string
type Jump string

const (
	DEST_M   Dest = "M"
	DEST_D   Dest = "D"
	DEST_MD  Dest = "MD"
	DEST_A   Dest = "A"
	DEST_AM  Dest = "AM"
	DEST_AD  Dest = "AD"
	DEST_AMD Dest = "AMD"
)

const (
	COMP_0         Comp = "0"
	COMP_1         Comp = "1"
	COMP_MINUS_1   Comp = "-1"
	COMP_D         Comp = "D"
	COMP_A         Comp = "A"
	COMP_NOT_D     Comp = "!D"
	COMP_NOT_A     Comp = "!A"
	COMP_MINUS_A   Comp = "-A"
	COMP_MINUS_D   Comp = "-D"
	COMP_D_ADD_1   Comp = "D+1"
	COMP_A_ADD_1   Comp = "A+1"
	COMP_D_MINUS_1 Comp = "D-1"
	COMP_A_MINUS_1 Comp = "A-1"
	COMP_D_ADD_A   Comp = "D+A"
	COMP_D_MINUS_A Comp = "D-A"
	COMP_A_MINUS_D Comp = "A-D"
	COMP_D_AND_A   Comp = "D&A"
	COMP_D_OR_A    Comp = "D|A"
	COMP_M         Comp = "M"
	COMP_MINUS_M   Comp = "-M"
	COMP_NOT_M     Comp = "!M"
	COMP_M_ADD_1   Comp = "M+1"
	COMP_M_MINUS_1 Comp = "M-1"
	COMP_D_ADD_M   Comp = "D+M"
	COMP_D_MINUS_M Comp = "D-M"
	COMP_M_MINUS_D Comp = "M-D"
	COMP_D_AND_M   Comp = "D&M"
	COMP_D_OR_M    Comp = "D|M"
)

const (
	JUMP_GREATER_THAN  Jump = "JGT"
	JUMP_EQUAL         Jump = "JEQ"
	JUMP_GREATER_EQUAL Jump = "JGE"
	JUMP_LOWER_THAN    Jump = "JLT"
	JUMP_NOT_EQUAL     Jump = "JNE"
	JUMP_LOWER_EQUAL   Jump = "JLE"
	JUMP               Jump = "JMP"
)

func mapDestToBinary(dest Dest) (string, error) {
	mp := map[Dest]string{
		DEST_M:   "001",
		DEST_D:   "010",
		DEST_MD:  "011",
		DEST_A:   "100",
		DEST_AM:  "101",
		DEST_AD:  "110",
		DEST_AMD: "111",
	}

	b, ok := mp[dest]
	if !ok {
		return "", ErrNoSuchDest
	}

	return b, nil
}

func mapCompToBinary(comp Comp) (string, error) {
	mp := map[Comp]string{
		COMP_0:         "0101010",
		COMP_1:         "0111111",
		COMP_MINUS_1:   "0111010",
		COMP_D:         "0001100",
		COMP_A:         "0110000",
		COMP_NOT_D:     "0001101",
		COMP_NOT_A:     "0110001",
		COMP_MINUS_D:   "0001111",
		COMP_MINUS_A:   "0110011",
		COMP_D_ADD_1:   "0011111",
		COMP_A_ADD_1:   "0110111",
		COMP_D_MINUS_1: "0001110",
		COMP_A_MINUS_1: "0110010",
		COMP_D_ADD_A:   "0000010",
		COMP_D_MINUS_A: "0010011",
		COMP_A_MINUS_D: "0000111",
		COMP_D_AND_A:   "0000000",
		COMP_D_OR_A:    "0010101",
		COMP_M:         "1110000",
		COMP_NOT_M:     "1110001",
		COMP_MINUS_M:   "1110011",
		COMP_M_ADD_1:   "1110111",
		COMP_M_MINUS_1: "1110010",
		COMP_D_ADD_M:   "1000010",
		COMP_D_MINUS_M: "1010011",
		COMP_M_MINUS_D: "1000111",
		COMP_D_AND_M:   "1000000",
		COMP_D_OR_M:    "1010101",
	}

	b, ok := mp[comp]
	if !ok {
		return "", ErrNoSuchComp
	}

	return b, nil
}

func mapJumpToBinary(jump Jump) (string, error) {
	mp := map[Jump]string{
		JUMP_GREATER_THAN:  "001",
		JUMP_EQUAL:         "010",
		JUMP_GREATER_EQUAL: "011",
		JUMP_LOWER_THAN:    "100",
		JUMP_NOT_EQUAL:     "101",
		JUMP_LOWER_EQUAL:   "110",
		JUMP:               "111",
	}

	b, ok := mp[jump]
	if !ok {
		return "", ErrNoSuchJump
	}

	return b, nil
}
