package codewriter

type Command interface {
	convert() ([]byte, error)
}
