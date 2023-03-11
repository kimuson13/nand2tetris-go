package parser

import (
	"assembler/code"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NEW_LINE = "\r\n"

type commandType string

var (
	ErrEmptyString       = errors.New("empty string")
	ErrEqualCountTooMany = errors.New("equal is too many")
	ErrSemiColonTooMany  = errors.New("semicolon is too many")
	ErrInvalidCommand    = errors.New("invalid command")
)

type Parser struct {
	commands   []string // アセンブリファイルを改行ごとにする
	currentIdx int      // 現在の実行行数
}

func Parse(fileName string) ([]code.Command, error) {
	parser, err := prepare(fileName)
	results := make([]code.Command, 0, len(parser.commands))

	if err != nil {
		return results, fmt.Errorf("parse error: %w", err)
	}

	for parser.hasMoreCommand() {
		command, err := parser.commandType()
		if err != nil {
			return results, fmt.Errorf("parse error: %w", err)
		}

		res := command.parse()
		results = append(results, res)

		parser.advance()
	}

	return results, nil
}

func prepare(fileName string) (Parser, error) {
	b, err := os.ReadFile(fileName)
	var p Parser
	if err != nil {
		return p, fmt.Errorf("prepare error: %w", err)
	}

	input := strings.ReplaceAll(string(b), " ", "")
	commands := strings.Split(input, NEW_LINE)
	p = Parser{commands: commands, currentIdx: 0}

	return p, nil
}

func (p *Parser) hasMoreCommand() bool {
	return len(p.commands) > p.currentIdx
}

func (p *Parser) advance() {
	p.currentIdx++
}

func (p *Parser) commandType() (Command, error) {
	currentCommand := p.commands[p.currentIdx]

	isEmptyString := currentCommand == ""
	if isEmptyString {
		return nil, ErrEmptyString
	}

	head := currentCommand[0]

	// TODO: 切り分ける
	isACommand := head == '@'
	if isACommand {
		val := string(currentCommand[0:])
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		command := &ACommand{value: i, symbol: val}
		return command, nil
	}

	// TODO: 切り分ける
	isLCommand := head == '('
	if isLCommand {
		val := string(currentCommand[0 : len(currentCommand)-1])

		command := &LCommand{value: 0, symbol: val}
		return command, nil
	}

	// TODO: 切り分ける
	isIncludeEq := strings.Contains(currentCommand, "=")
	eqCnt := strings.Count(currentCommand, "=")
	eqIdx := strings.Index(currentCommand, "=")
	if eqCnt > 1 {
		return nil, ErrEqualCountTooMany
	}

	isIncludeSemiColon := strings.Contains(currentCommand, ";")
	semiColonCnt := strings.Count(currentCommand, ";")
	semiColonIdx := strings.Index(currentCommand, ";")
	if semiColonCnt > 1 {
		return nil, ErrSemiColonTooMany
	}

	// dest=comp;jump
	isCompDestJump := isIncludeEq && isIncludeSemiColon && eqIdx < semiColonIdx
	if isCompDestJump {
		dest := string(currentCommand[0:eqIdx])
		comp := string(currentCommand[eqIdx:semiColonIdx])
		jump := string(currentCommand[semiColonIdx:])

		command := &CCommand{dest: dest, comp: comp, jump: jump}
		return command, nil
	}

	// dest=comp
	isDestJump := isIncludeEq && !isIncludeSemiColon
	if isDestJump {
		dest := string(currentCommand[0:eqIdx])
		comp := string(currentCommand[eqIdx:])

		command := &CCommand{dest: dest, comp: comp, jump: ""}
		return command, nil
	}

	// comp;jump
	isCompJump := !isIncludeEq && isIncludeSemiColon
	if isCompJump {
		comp := string(currentCommand[0:semiColonIdx])
		jump := string(currentCommand[semiColonIdx:])

		command := &CCommand{dest: "", comp: comp, jump: jump}
		return command, nil
	}

	return nil, ErrInvalidCommand
}
