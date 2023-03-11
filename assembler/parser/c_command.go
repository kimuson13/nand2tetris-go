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

func (c *CCommand) parse() code.Command {
	return nil
}

func toCCommand(raw string) (*CCommand, error) {
	const TO_CCOMAND_ERR = "toCCommand error: %w"

	stmt, err := genCCommandStmt(raw)
	if err != nil {
		return nil, fmt.Errorf(TO_CCOMAND_ERR, err)
	}

	cCommad, err := stmt.toCCommand()
	if err != nil {
		return nil, fmt.Errorf(TO_CCOMAND_ERR, err)
	}

	return cCommad, nil
}

func isCComand(raw string) bool {
	eqCnt := strings.Count(raw, "=")
	if eqCnt > 1 {
		return false
	}

	semiColonCnt := strings.Count(raw, ";")
	if semiColonCnt > 1 {
		return false
	}

	return true
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
		return toDestCompJump(c.raw, c.eqPos, c.semiColonPos)
	case DEST_COMP:
		return toDestComp(c.raw, c.eqPos)
	case COMP_JUMP:
		return toCompJump(c.raw, c.semiColonPos)
	}

	return nil, ErrCanNotParseCCommand
}

func toDestCompJump(raw string, eqIdx, semiColonIdx int) (*CCommand, error) {
	dest := string(raw[0:eqIdx])
	comp := string(raw[eqIdx:semiColonIdx])
	jump := string(raw[semiColonIdx:])

	command := &CCommand{dest: dest, comp: comp, jump: jump}
	return command, nil
}

func toDestComp(raw string, eqIdx int) (*CCommand, error) {
	dest := string(raw[0:eqIdx])
	comp := string(raw[eqIdx:])

	command := &CCommand{dest: dest, comp: comp, jump: ""}
	return command, nil
}

func toCompJump(raw string, semiColonIdx int) (*CCommand, error) {
	comp := string(raw[0:semiColonIdx])
	jump := string(raw[semiColonIdx:])

	command := &CCommand{dest: "", comp: comp, jump: jump}
	return command, nil
}
