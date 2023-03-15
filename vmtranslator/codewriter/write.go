package codewriter

import (
	"fmt"
	"os"
)

type CodeWriter struct {
	file     *os.File
	commands []Command
}

func New(commands []Command, fileName string) (CodeWriter, error) {
	var codeWriter CodeWriter
	asmFileName := fmt.Sprintf("%s.asm", fileName)

	f, err := os.Create(asmFileName)
	if err != nil {
		return codeWriter, fmt.Errorf("code writer new error: %w", err)
	}

	codeWriter.commands = commands
	codeWriter.file = f

	return codeWriter, nil
}

func (c CodeWriter) Write() error {
	for _, command := range c.commands {
		b, err := command.convert()
		if err != nil {
			return fmt.Errorf("write error: %w", err)
		}

		if _, err := c.file.Write(b); err != nil {
			return fmt.Errorf("write error: %w", err)
		}
	}

	if err := c.file.Close(); err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	return nil
}
