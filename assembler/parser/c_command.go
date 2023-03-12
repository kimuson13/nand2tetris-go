package parser

import (
	"assembler/code"
	"errors"
	"fmt"
	"strings"
)

type CCommandKind int

const (
	DEST_COMP_JUMP CCommandKind = iota
	DEST_COMP
	COMP_JUMP
)

var (
	ErrCanNotParseCCommand = errors.New("can not parse to CCommand")
	ErrNoSuchComp          = errors.New("no such a comp kind")
	ErrNoSuchDest          = errors.New("no such a dest kind")
	ErrNoSuchJump          = errors.New("no such a jump kind")
)

type CCommand struct {
	dest string
	comp string
	jump string
}

type CCommandStmt struct {
	raw          string
	kind         CCommandKind
	eqPos        int
	semiColonPos int
}

func isCCommand(raw string) bool {
	eqCnt := strings.Count(raw, "=")
	if eqCnt > 1 {
		return false
	}
	eqIdx := strings.Index(raw, "=")

	semiColonCnt := strings.Count(raw, ";")
	if semiColonCnt > 1 {
		return false
	}
	semiColonIdx := strings.Index(raw, ";")

	if eqCnt == 1 && semiColonCnt == 1 {
		return eqIdx < semiColonIdx
	}

	return true
}

func (c *CCommand) parse() (code.Command, error) {
	dest, err := mapCodeDest(c.dest)
	if err != nil {
		return nil, err
	}

	comp, err := mapCodeComp(c.comp)
	if err != nil {
		return nil, err
	}

	jump, err := mapCodeJump(c.jump)
	if err != nil {
		return nil, err
	}

	return &code.CCommand{
		Dest: dest,
		Comp: comp,
		Jump: jump,
	}, nil
}

func mapCodeDest(raw string) (*code.Dest, error) {
	mp := map[string]code.Dest{
		"M":   code.DEST_M,
		"D":   code.DEST_D,
		"MD":  code.DEST_MD,
		"A":   code.DEST_A,
		"AM":  code.DEST_AM,
		"AD":  code.DEST_AD,
		"AMD": code.DEST_AMD,
	}

	if raw == "" {
		return nil, nil
	}

	dest, ok := mp[raw]
	if !ok {
		return nil, ErrNoSuchDest
	}

	return &dest, nil
}

func mapCodeComp(raw string) (code.Comp, error) {
	mp := map[string]code.Comp{
		"0":   code.COMP_0,
		"1":   code.COMP_1,
		"-1":  code.COMP_MINUS_1,
		"D":   code.COMP_D,
		"A":   code.COMP_A,
		"!D":  code.COMP_NOT_D,
		"!A":  code.COMP_NOT_A,
		"-A":  code.COMP_MINUS_A,
		"-D":  code.COMP_MINUS_D,
		"D+1": code.COMP_D_ADD_1,
		"A+1": code.COMP_A_ADD_1,
		"D-1": code.COMP_D_MINUS_1,
		"A-1": code.COMP_A_MINUS_1,
		"D+A": code.COMP_D_ADD_A,
		"D-A": code.COMP_D_MINUS_A,
		"A-D": code.COMP_A_MINUS_D,
		"D&A": code.COMP_D_AND_A,
		"D|A": code.COMP_D_OR_A,
		"M":   code.COMP_M,
		"-M":  code.COMP_MINUS_M,
		"!M":  code.COMP_NOT_M,
		"M+1": code.COMP_M_ADD_1,
		"M-1": code.COMP_M_MINUS_1,
		"D+M": code.COMP_D_ADD_M,
		"D-M": code.COMP_D_MINUS_A,
		"M-D": code.COMP_M_MINUS_D,
		"D&M": code.COMP_D_AND_M,
		"D|M": code.COMP_D_OR_M,
	}

	comp, ok := mp[raw]
	if !ok {
		return "", ErrNoSuchComp
	}

	return comp, nil
}

func mapCodeJump(raw string) (*code.Jump, error) {
	mp := map[string]code.Jump{
		"JGT": code.JUMP_GREATER_THAN,
		"JEQ": code.JUMP_EQUAL,
		"JGE": code.JUMP_GREATER_EQUAL,
		"JLT": code.JUMP_LOWER_THAN,
		"JNE": code.JUMP_NOT_EQUAL,
		"JLE": code.JUMP_LOWER_EQUAL,
		"JMP": code.JUMP,
	}

	if raw == "" {
		return nil, nil
	}

	jump, ok := mp[raw]
	if !ok {
		return nil, ErrNoSuchJump
	}

	return &jump, nil
}

func toCCommand(raw string) (*CCommand, error) {
	const TO_CCOMMAND_ERR = "toCCommand error: %w"

	stmt, err := genCCommandStmt(raw)
	if err != nil {
		return nil, fmt.Errorf(TO_CCOMMAND_ERR, err)
	}

	cCommad, err := stmt.toCCommand()
	if err != nil {
		return nil, fmt.Errorf(TO_CCOMMAND_ERR, err)
	}

	return cCommad, nil
}

func genCCommandStmt(raw string) (CCommandStmt, error) {
	isIncludeEq := strings.Contains(raw, "=")
	eqIdx := strings.Index(raw, "=")

	isIncludeSemiColon := strings.Contains(raw, ";")
	semiColonIdx := strings.Index(raw, ";")
	res := CCommandStmt{raw: raw, eqPos: eqIdx, semiColonPos: semiColonIdx}

	isCompDestJump := isIncludeEq && isIncludeSemiColon && eqIdx < semiColonIdx
	if isCompDestJump {
		res.kind = DEST_COMP_JUMP
		return res, nil
	}

	isDestComp := isIncludeEq && !isIncludeSemiColon
	if isDestComp {
		res.kind = DEST_COMP
		return res, nil
	}

	isCompJump := !isIncludeEq && isIncludeSemiColon
	if isCompJump {
		res.kind = COMP_JUMP
		return res, nil
	}

	return res, ErrCanNotParseCCommand
}

func (c *CCommandStmt) toCCommand() (*CCommand, error) {
	switch c.kind {
	case DEST_COMP_JUMP:
		return c.toDestCompJump()
	case DEST_COMP:
		return c.toDestComp()
	case COMP_JUMP:
		return c.toCompJump()
	}

	return nil, ErrCanNotParseCCommand
}

func (c *CCommandStmt) toDestCompJump() (*CCommand, error) {
	dest := string(c.raw[:c.eqPos])
	comp := string(c.raw[c.eqPos+1 : c.semiColonPos])
	jump := string(c.raw[c.semiColonPos+1:])

	command := &CCommand{dest: dest, comp: comp, jump: jump}
	return command, nil
}

func (c *CCommandStmt) toDestComp() (*CCommand, error) {
	dest := string(c.raw[:c.eqPos])
	comp := string(c.raw[c.eqPos+1:])

	command := &CCommand{dest: dest, comp: comp, jump: ""}
	return command, nil
}

func (c *CCommandStmt) toCompJump() (*CCommand, error) {
	comp := string(c.raw[:c.semiColonPos])
	jump := string(c.raw[c.semiColonPos+1:])

	command := &CCommand{dest: "", comp: comp, jump: jump}
	return command, nil
}
