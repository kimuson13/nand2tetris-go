package code

import (
	"fmt"
)

type Command interface {
	Convert() (string, error)
}

type CCommand struct {
	Dest *Dest
	Comp Comp
	Jump *Jump
}

func (c *CCommand) Convert() (string, error) {
	dest, err := c.dest()
	if err != nil {
		return "", fmt.Errorf("c command convert error: %w", err)
	}

	comp, err := c.comp()
	if err != nil {
		return "", fmt.Errorf("c command convert error: %w", err)
	}

	jump, err := c.jump()
	if err != nil {
		return "", fmt.Errorf("c command convert error: %w", err)
	}

	return fmt.Sprintf("111%s%s%s", comp, dest, jump), nil
}

func (c *CCommand) dest() (string, error) {
	if c.Dest == nil {
		return "000", nil
	}

	b, err := mapDestToBinary(*c.Dest)
	if err != nil {
		return "", fmt.Errorf("convert dest error: %w", err)
	}

	return b, nil
}

func (c *CCommand) comp() (string, error) {
	b, err := mapCompToBinary(c.Comp)
	if err != nil {
		return "", fmt.Errorf("convert comp error: %w", err)
	}

	return b, nil
}

func (c *CCommand) jump() (string, error) {
	if c.Jump == nil {
		return "000", nil
	}

	b, err := mapJumpToBinary(*c.Jump)
	if err != nil {
		return "", fmt.Errorf("convert jump error: %w", err)
	}

	return b, nil
}

type ACommand struct {
	Address int
	Symbol  string
}

func (a *ACommand) Convert() (string, error) {
	return fmt.Sprintf("0%015b", a.Address), nil
}

type LCommand struct {
	Value  int
	Symbol string
}

func (l *LCommand) Convert() (string, error) {
	return "", nil
}
