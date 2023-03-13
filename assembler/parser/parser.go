package parser

import (
	"assembler/code"
	"assembler/symtable"
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
	ErrInvalidSyntax     = errors.New("invalid syntax")
	ErrTooManyCommentLit = errors.New("too many comment literal")
)

type Parser struct {
	commands    []string // アセンブリファイルを改行ごとにする
	currentIdx  int      // 現在の実行行数
	symbolTable symtable.SymTable
}

func New(path string) (Parser, error) {
	b, err := os.ReadFile(path)
	var p Parser
	if err != nil {
		return p, fmt.Errorf("prepare error: %w", err)
	}

	lines := strings.Split(string(b), NEW_LINE)
	commands, err := getCommands(lines)
	if err != nil {
		return p, fmt.Errorf("prepare error: %w", err)
	}

	p = Parser{commands: commands, currentIdx: 0}

	return p, nil
}

func getCommands(lines []string) ([]string, error) {
	commands := make([]string, 0, len(lines))
	for _, line := range lines {
		command, err := getCommand(line)
		if err != nil {
			return commands, fmt.Errorf("prepare error: %w", err)
		}

		if command != "" {
			commands = append(commands, command)
		}
	}

	return commands, nil
}

func getCommand(raw string) (string, error) {
	line := strings.TrimSpace(raw)
	if line == "" {
		return "", nil
	}

	commentCnt := strings.Count(line, "//")
	if commentCnt > 1 {
		return "", ErrTooManyCommentLit
	}

	if string(line[0:2]) == "//" {
		return "", nil
	}

	isInlineComment := strings.Contains(line, "//")
	if isInlineComment {
		commentIdx := strings.Index(line, "//")
		line = strings.TrimSpace(string(line[:commentIdx]))
	}
	isIncludeSpaceOrTab := strings.Contains(line, " ") || strings.Contains(line, "\t")
	if isIncludeSpaceOrTab {
		return "", ErrInvalidSyntax
	}

	return line, nil
}

func (p *Parser) SynbolicLink() error {
	return nil
}

func (p *Parser) Parse() ([]code.Command, error) {
	results := make([]code.Command, 0, len(p.commands))

	for p.hasMoreCommand() {
		command, err := p.commandType()
		if err != nil {
			return results, fmt.Errorf("parse error: %w", err)
		}

		res, err := command.parse()
		if err != nil {
			return results, fmt.Errorf("parse error: %w", err)
		}
		results = append(results, res)

		p.advance()
	}

	return results, nil
}

func (p *Parser) hasMoreCommand() bool {
	return len(p.commands) > p.currentIdx
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

func (p *Parser) advance() {
	p.currentIdx++
}

func (p *Parser) resetCurrentIdx() {
	p.currentIdx = 0
}
