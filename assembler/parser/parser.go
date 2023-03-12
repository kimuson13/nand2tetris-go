package parser

import (
	"assembler/code"
	"errors"
	"fmt"
	"os"
	"strings"
)

const NEW_LINE = "\r\n"

var (
	ErrEmptyString       = errors.New("empty string")
	ErrEqualCountTooMany = errors.New("equal is too many")
	ErrSemiColonTooMany  = errors.New("semicolon is too many")
	ErrInvalidCommand    = errors.New("invalid command")
	ErrTooManyCommentLit = errors.New("too many comment literal")
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

		res, err := command.parse()
		if err != nil {
			return results, fmt.Errorf("parse error: %w", err)
		}
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
	lines := strings.Split(input, NEW_LINE)
	commands := make([]string, 0, len(lines))
	for _, line := range lines {
		command, err := getCommand(line)
		if err != nil {
			return p, fmt.Errorf("prepare error: %w", err)
		}

		if command != "" {
			commands = append(commands, command)
		}
	}

	p = Parser{commands: commands, currentIdx: 0}

	return p, nil
}

func getCommand(raw string) (string, error) {
	commentCnt := strings.Count(raw, "//")
	if commentCnt > 1 {
		return "", ErrTooManyCommentLit
	}

	if string(raw[0:2]) == "//" {
		return "", nil
	}

	isInlineComment := strings.Contains(raw, "//")
	if isInlineComment {
		commentIdx := strings.Index(raw, "//")
		return string(raw[:commentIdx]), nil
	}

	return raw, nil
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

	if isACommand(currentCommand) {
		return toACommand(currentCommand)
	}

	if isLCommand(currentCommand) {
		return toLCommand(currentCommand)
	}

	if isCCommand(currentCommand) {
		return toCCommand(currentCommand)
	}

	return nil, ErrInvalidCommand
}
