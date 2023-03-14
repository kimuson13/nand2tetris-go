package parser

import (
	"errors"
	"fmt"
	"strings"
)

type Parser struct {
	currentIdx     int
	currentCommand []string
	commands       []string
}

var (
	ErrTooManyCommentLiteral = errors.New("too many comment literal")
)

func New(raw string) (Parser, error) {
	rawLines := strings.Split(raw, "\n")
	commands := make([]string, 0, len(rawLines))
	p := Parser{}
	for _, line := range rawLines {
		trimLine := strings.TrimSpace(line)

		if isNotCommand(trimLine) {
			continue
		}

		isIncludeComment := strings.Contains(trimLine, "//")
		commentLitCnt := strings.Count(trimLine, "//")
		if commentLitCnt > 1 {
			return p, fmt.Errorf("parser new error: %w", ErrTooManyCommentLiteral)
		}

		if isIncludeComment {
			commentIdx := strings.Index(trimLine, "//")
			trimLine = trimLine[:commentIdx]
		}

		commands = append(commands, trimLine)
	}

	p.commands = commands
	firstCommand := strings.Split(p.commands[0], " ")
	p.currentCommand = firstCommand

	return p, nil
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
