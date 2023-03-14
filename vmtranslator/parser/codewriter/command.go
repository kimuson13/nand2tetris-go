package codewriter

type Command interface {
	convert() ([]byte, error)
}

type Arithmetic struct {
}

func (a Arithmetic) convert() ([]byte, error) {
	return nil, nil
}
