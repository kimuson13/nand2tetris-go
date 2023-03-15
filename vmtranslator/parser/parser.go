package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"vmtranslator/codewriter"
)

const NEW_LINE = "\n"

var (
	ErrInvalidCommand        = errors.New("invalid command")
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

func (p Parser) Parse() ([]codewriter.Command, error) {
	results := make([]codewriter.Command, 0, len(p.commands))

	for p.hasMoreCommand() {
		cType := p.commandType()
		command, err := p.parse(cType)
		if err != nil {
			return results, fmt.Errorf("Parse error: %w", err)
		}

		if command != nil {
			results = append(results, command)
		}

		p.advance()
	}

	return results, nil
}

func (p Parser) hasMoreCommand() bool {
	return p.currentIdx < len(p.commands)
}

func (p *Parser) advance() {
	p.currentIdx++
	if p.hasMoreCommand() {
		nextCommand := p.commands[p.currentIdx]
		p.currentCommand = strings.Split(nextCommand, " ")
	}
}

func (p Parser) commandType() command {
	head := commandType(p.currentCommand[0])
	switch head {
	case ADD, SUB, NEGATIVE, EQUAL, GREATER_THAN, LOWER_THAN, AND, OR, NOT:
		return C_ARITHMETIC
	case PUSH:
		return C_PUSH
	}

	return INVALID
}

func (p Parser) parse(c command) (codewriter.Command, error) {
	switch c {
	case C_ARITHMETIC:
		command, err := p.parseArithmetic()
		if err != nil {
			return nil, fmt.Errorf("parse error: %w", err)
		}
		return command, nil
	case C_PUSH:
		command, err := p.parsePush()
		if err != nil {
			return nil, fmt.Errorf("parse error: %w", err)
		}
		return command, nil
	}

	return nil, fmt.Errorf("parse error: %w", ErrNoSuchACommandType)
}

func (p Parser) parseArithmetic() (codewriter.Arithmetic, error) {
	arithmetic := codewriter.Arithmetic{}
	cType, err := p.arg1(C_ARITHMETIC)
	if err != nil {
		return arithmetic, fmt.Errorf("arithmetci error : %w", err)
	}

	kind, err := mapCommandTypeToArithmeticKind(commandType(cType))
	if err != nil {
		return arithmetic, fmt.Errorf("arithmetic error: %w", err)
	}

	arithmetic.Kind = kind
	return arithmetic, nil
}

func (p Parser) parsePush() (codewriter.Push, error) {
	var push codewriter.Push
	arg1, err := p.arg1(C_PUSH)
	if err != nil {
		return push, fmt.Errorf("push error: %w", err)
	}

	segment, err := mapSegment(arg1)
	if err != nil {
		return push, fmt.Errorf("push error: %w", err)
	}

	index, err := p.arg2(C_PUSH)
	if err != nil {
		return push, fmt.Errorf("push error: %w", err)
	}

	push.Segment = segment
	push.Index = index

	return push, nil
}

func (p Parser) arg1(c command) (string, error) {
	if c == C_ARITHMETIC {
		return p.currentCommand[0], nil
	}

	return p.currentCommand[1], nil
}

func (p Parser) arg2(c command) (int, error) {
	if c == C_PUSH {
		arg2, err := strconv.Atoi(p.currentCommand[2])
		if err != nil {
			return 0, fmt.Errorf("arg2 error: %w", err)
		}

		return arg2, nil
	}

	return 0, fmt.Errorf("arg2 error: %w", ErrInvalidCommand)
}
