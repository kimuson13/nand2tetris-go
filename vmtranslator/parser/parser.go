package parser

import (
	"errors"
	"fmt"
	"strings"
)

const NEW_LINE = "\n"

var (
	ErrNoSuchACommandType    = errors.New("no such a command type")
	ErrTooManyCommentLiteral = errors.New("too many comment literal")
)

type Parser struct {
	currentIdx     int
	currentCommand []string
	commands       []string
}

func New(raw string) (Parser, error) {
	p := Parser{}

	commands, err := getCommands(raw)
	if err != nil {
		return p, fmt.Errorf("parser new error: %w", err)
	}

	p.commands = commands
	firstCommand := strings.Split(p.commands[0], " ")
	p.currentCommand = firstCommand

	return p, nil
}

func getCommands(raw string) ([]string, error) {
	rawLines := strings.Split(raw, NEW_LINE)
	commands := make([]string, 0, len(rawLines))
	for _, line := range rawLines {
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

func getCommand(line string) (string, error) {
	spaceTrimedLine := strings.TrimSpace(line)

	if isNotCommand(spaceTrimedLine) {
		return "", nil
	}

	command, err := trimComment(spaceTrimedLine)
	if err != nil {
		return "", fmt.Errorf("getCommand error: %w", err)
	}

	return command, nil
}

func isNotCommand(val string) bool {
	isEmpty := val == ""
	if isEmpty {
		return true
	}

	isCommentLine := strings.HasPrefix(val, "//")
	if isCommentLine {
		return true
	}

	return false
}

func trimComment(line string) (string, error) {
	isIncludeComment := strings.Contains(line, "//")
	commentLitCnt := strings.Count(line, "//")
	if commentLitCnt > 1 {
		return "", fmt.Errorf("trimComment error: %w", ErrTooManyCommentLiteral)
	}

	if isIncludeComment {
		commentIdx := strings.Index(line, "//")
		line = line[:commentIdx]
	}

	return strings.TrimSpace(line), nil
}

func (p Parser) hasMoreCommand() bool {
	return p.currentIdx < len(p.commands)
}

func (p *Parser) advance() {
	p.currentIdx++
	nextCommand := p.commands[p.currentIdx]
	p.currentCommand = strings.Split(nextCommand, " ")
}

func (p Parser) commandType() command {
	head := commandType(p.currentCommand[0])
	switch head {
	case ADD:
		return C_ARITHMETIC
	case PUSH:
		return C_PUSH
	}

	return INVALID
}
