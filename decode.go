package main;

import(
	"strings"
)

func decode(  inst string ) Instruction{
	i := strings.Split(inst, " ")
	var newInst Instruction
	if(i[0] == "ADD"){
		newInst.Opcode = ADD
	} else if(i[0] == "SUB"){
		newInst.Opcode = SUB
	} else if(i[0] == "MUL"){
		newInst.Opcode = MUL
	}

	newInst.DestinationReg = i[1]
	newInst.SourceReg1 = i[2]
	newInst.SourceReg2 = i[3]

	return newInst
}


/*
type Instruction struct {
	Opcode int // Operation (e.g., ADD, SUB, MUL)
	DestinationReg    string // Destination register
	SourceReg1    Register // Source register 1
	SourceReg2    Register // Source register 2
	InstructionType string    // Execution cycles
} 
*/