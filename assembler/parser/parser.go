package parser

import (
	"assembler/code"
	"fmt"
	"os"
	"strings"
)

const NEW_LINE = "\r\n"

type Parser struct {
	commands   []string // アセンブリファイルを改行ごとにする
	currentIdx int      // 現在の実行行数
}

func Parse(fileName string) ([]code.Command, error) {
	parser, err := prepare(fileName)
	if err != nil {
		return []code.Command{}, fmt.Errorf("parse error: %w", err)
	}

	results := make([]code.Command, 0, len(parser.commands))
	for parser.hasMoreCommand() {
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
