package codewriter

type Pop struct {
	FileName string
	Segment  Segment
	Index    int
}

func (p Pop) convert() ([]byte, error) {
	return nil, nil
}
