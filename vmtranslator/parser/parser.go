package parser

import (
	"errors"
	"fmt"
	"strings"
)

const NEW_LINE = "\n"

var (
	ErrTooManyCommentLiteral = errors.New("too many comment literal")
)

type Parser struct {
	currentIdx     int
	currentCommand []string
	commands       []string
}

func New(raw string) (Parser, error) {
	rawLines := strings.Split(raw, NEW_LINE)
	commands := make([]string, 0, len(rawLines))
	p := Parser{}
	for _, line := range rawLines {
		spaceTrimedLine := strings.TrimSpace(line)

		if isNotCommand(spaceTrimedLine) {
			continue
		}

		command, err := trimComment(spaceTrimedLine)
		if err != nil {
			return p, fmt.Errorf("parser new error: %w", err)
		}

		commands = append(commands, command)
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

func trimComment(line string) (string, error) {
	isIncludeComment := strings.Contains(line, "//")
	commentLitCnt := strings.Count(line, "//")
	if commentLitCnt > 1 {
		return "", fmt.Errorf("trimComment error: %w", ErrTooManyCommentLiteral)
	}

	if isIncludeComment {
		commentIdx := strings.Index(line, "//")
		return line[:commentIdx], nil
	}

	return line, nil
}
