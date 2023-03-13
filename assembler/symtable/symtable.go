package symtable

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

func New() *SymTable {
	symMap := genDefinedSymbol()

	return &SymTable{
		mp: symMap,
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
