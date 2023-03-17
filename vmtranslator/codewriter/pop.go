package codewriter

type Pop struct {
	Segment Segment
	Index   int
}

func (p Pop) convert() ([]byte, error) {
	return nil, nil
}
