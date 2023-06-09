package parser

import (
	"assembler/code"
	"assembler/symtable"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NEW_LINE = "\n"

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
			return commands, fmt.Errorf("getCommands error: %w", err)
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
		return "", fmt.Errorf("getCommand error: %w, raw=%s", ErrTooManyCommentLit, raw)
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

func (p *Parser) SymbolicLink() error {
	symTable := symtable.New()
	p.symbolTable = symTable

	if err := p.linkLCommandSymbol(); err != nil {
		return fmt.Errorf("symbolic link error: %w", err)
	}

	if err := p.linkACommandSymbol(); err != nil {
		return fmt.Errorf("symbolic link error: %w", err)
	}

	return nil
}

func (p *Parser) linkLCommandSymbol() error {
	idx := 0
	for p.hasMoreCommand() {
		command, err := p.commandType()
		if err != nil {
			return fmt.Errorf("link LCommand error: %w", err)
		}

		lCommand, ok := command.(*lCommand)
		if ok {
			if err := p.addEntryToSymTable(lCommand.symbol, idx); err != nil {
				return fmt.Errorf("link LCommand error: %w", err)
			}
		} else {
			idx++
		}

		p.advance()
	}

	p.resetCurrentIdx()
	return nil
}

func (p *Parser) linkACommandSymbol() error {
	const MIN_SYMBOL_ADDRESS = 16
	aCommandSymbolCnt := 0
	for p.hasMoreCommand() {
		command, err := p.commandType()
		if err != nil {
			return fmt.Errorf("link ACommand error: %w", err)
		}

		aCommand, ok := command.(*aCommand)
		if ok {
			symbol := aCommand.symbol
			address := aCommandSymbolCnt + MIN_SYMBOL_ADDRESS
			if !p.symbolTable.Contains(symbol) && symbol != "" {
				if err := p.symbolTable.AddEntry(symbol, address); err != nil {
					return fmt.Errorf("link ACommand error: %w", err)
				}
				aCommandSymbolCnt++
			}
		}

		p.advance()
	}

	p.resetCurrentIdx()
	return nil
}

func (p *Parser) addEntryToSymTable(symbol string, address int) error {
	if p.symbolTable.Contains(symbol) {
		return nil
	}

	if err := p.symbolTable.AddEntry(symbol, address); err != nil {
		return err
	}

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
		return nil, fmt.Errorf("commandType error: %w", ErrEmptyString)
	}

	if isACommand(currentCommand) {
		command, err := p.toACommand(currentCommand)
		if err != nil {
			return nil, fmt.Errorf("commandType error: %w", err)
		}
		return command, nil
	}

	if isLCommand(currentCommand) {
		command, err := p.toLCommand(currentCommand)
		if err != nil {
			return nil, fmt.Errorf("commandType error: %w", err)
		}
		return command, nil
	}

	if isCCommand(currentCommand) {
		command, err := p.toCCommand(currentCommand)
		if err != nil {
			return nil, fmt.Errorf("commandType error: %w", err)
		}
		return command, nil
	}

	return nil, fmt.Errorf("commandType error: %w", ErrInvalidCommand)
}

func (p *Parser) toACommand(raw string) (*aCommand, error) {
	val := string(raw[1:])
	if isNumeric(val) {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("toACommand error: %w", err)
		}
		return &aCommand{address: i, symbol: ""}, nil
	}

	if p.symbolTable.Contains(val) {
		address, err := p.symbolTable.GetAddress(val)
		if err != nil {
			return nil, fmt.Errorf("toACommand error: %w", err)
		}
		return &aCommand{address: address, symbol: val}, nil
	}

	return &aCommand{symbol: val}, nil
}

func (p *Parser) toLCommand(raw string) (*lCommand, error) {
	val := string(raw[1 : len(raw)-1])

	return &lCommand{symbol: val}, nil
}

func (p *Parser) toCCommand(raw string) (*cCommand, error) {
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

func isNumeric(val string) bool {
	_, err := strconv.Atoi(val)
	return err == nil
}

func (p *Parser) advance() {
	p.currentIdx++
}

func (p *Parser) resetCurrentIdx() {
	p.currentIdx = 0
}
