package process

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"vmtranslator/codewriter"
	"vmtranslator/parser"
)

var (
	ErrInvalidArgs          = errors.New("need 1 arg")
	ErrInvalidFileExtension = errors.New("extension need .vm")
)

func Run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("process error: %w", ErrInvalidArgs)
	}

	dir, base := filepath.Split(args[0])
	ext := filepath.Ext(args[0])
	if ext != ".vm" {
		return fmt.Errorf("process error: %w", ErrInvalidFileExtension)
	}
	extIdx := strings.Index(base, ext)
	fileName := base[:extIdx]

	b, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("process error: %w", err)
	}

	p, err := parser.New(string(b), fileName)
	if err != nil {
		return fmt.Errorf("process error: %w", err)
	}

	commands, err := p.Parse()
	if err != nil {
		return fmt.Errorf("procee error: %w", err)
	}

	codeWriter, err := codewriter.New(commands, filepath.Join(dir, fileName))
	if err != nil {
		return fmt.Errorf("process error: %w", err)
	}

	if err := codeWriter.Write(); err != nil {
		return fmt.Errorf("process error: %w", err)
	}

	return nil
}
