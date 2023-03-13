package symtable

import "errors"

type SymTable struct {
	mp symMap
}

type symMap map[string]int

const (
	SP     = "SP"
	LCL    = "LCL"
	ARG    = "ARG"
	THIS   = "THIS"
	THAT   = "THAT"
	R0     = "R0"
	R1     = "R1"
	R2     = "R2"
	R3     = "R3"
	R4     = "R4"
	R5     = "R5"
	R6     = "R6"
	R7     = "R7"
	R8     = "R8"
	R9     = "R9"
	R10    = "R10"
	R11    = "R11"
	R12    = "R12"
	R13    = "R13"
	R14    = "R14"
	R15    = "R15"
	SCREEN = "SCREEN"
	KBD    = "KBD"
)

var (
	ErrDuplicateKeyEntry = errors.New("insert key is duplicated")
	ErrNoSuchAKey        = errors.New("no such a key")
)

func New() SymTable {
	return SymTable{
		mp: genDefinedSymbol(),
	}
}

func genDefinedSymbol() symMap {
	mp := map[string]int{
		SP:     0,
		LCL:    1,
		ARG:    2,
		THIS:   3,
		THAT:   4,
		R0:     0,
		R1:     1,
		R2:     2,
		R3:     3,
		R4:     4,
		R5:     5,
		R6:     6,
		R7:     7,
		R8:     8,
		R9:     9,
		R10:    10,
		R11:    11,
		R12:    12,
		R13:    13,
		R14:    14,
		R15:    15,
		SCREEN: 16348,
		KBD:    24578,
	}

	return mp
}

func (s *SymTable) AddEntry(symbol string, address int) error {
	if s.Contains(symbol) {
		return ErrDuplicateKeyEntry
	}

	s.mp[symbol] = address

	return nil
}

func (s *SymTable) GetAddress(symbol string) (int, error) {
	val, ok := s.mp[symbol]
	if !ok {
		return 0, ErrNoSuchAKey
	}

	return val, nil
}

func (s *SymTable) Contains(symbol string) bool {
	_, ok := s.mp[symbol]
	return ok
}
