package codewriter

type Push struct {
	Segment Segment
	Index   int
}

func (p Push) convert() ([]byte, error) {
	return nil, nil
}
