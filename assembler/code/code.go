package code

type Command interface {
	Convert() string
}

type CCommand struct {
	Dest *Dest
	Comp Comp
	Jump *Jump
}

func (c *CCommand) Convert() string {
	return ""
}

type ACommand struct {
	Value  int
	Symbol string
}

func (a *ACommand) Convert() string {
	return ""
}

type LCommand struct {
	Value  int
	Symbol string
}

func (l *LCommand) Convert() string {
	return ""
}

type Dest string
type Comp string
type Jump string

const (
	DEST_M   Dest = "M"
	DEST_D   Dest = "D"
	DEST_MD  Dest = "MD"
	DEST_A   Dest = "A"
	DEST_AM  Dest = "AM"
	DEST_AD  Dest = "AD"
	DEST_AMD Dest = "AMD"
)

const (
	COMP_0         Comp = "0"
	COMP_1         Comp = "1"
	COMP_MINUS_1   Comp = "-1"
	COMP_D         Comp = "D"
	COMP_A         Comp = "A"
	COMP_NOT_D     Comp = "!D"
	COMP_NOT_A     Comp = "!A"
	COMP_MINUS_A   Comp = "-A"
	COMP_MINUS_D   Comp = "-D"
	COMP_D_ADD_1   Comp = "D+1"
	COMP_A_ADD_1   Comp = "A+1"
	COMP_D_MINUS_1 Comp = "D-1"
	COMP_A_MINUS_1 Comp = "A-1"
	COMP_D_ADD_A   Comp = "D+A"
	COMP_D_MINUS_A Comp = "D-A"
	COMP_A_MINUS_D Comp = "A-D"
	COMP_D_AND_A   Comp = "D&A"
	COMP_D_OR_A    Comp = "D|A"
	COMP_M         Comp = "M"
	COMP_NOT_M     Comp = "!M"
	COMP_M_ADD_1   Comp = "M+1"
	COMP_D_ADD_M   Comp = "D+M"
	COMP_D_MINUS_M Comp = "D-M"
	COMP_D_AND_M   Comp = "D&M"
	COMP_D_OR_M    Comp = "D|M"
)

const (
	JUMP_GREATER_THAN  Jump = "JGT"
	JUMP_EQUAL         Jump = "JEQ"
	JUMP_GREATER_EQUAL Jump = "JGE"
	JUMP_LOWER_THAN    Jump = "JLT"
	JUMP_NOT_EQUAL     Jump = "JNE"
	JUMP_LOWER_EQUAL   Jump = "JLE"
	JUMP               Jump = "JMP"
)
