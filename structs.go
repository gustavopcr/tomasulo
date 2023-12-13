package main;

type Status int
type riscV int

const (
	ADD riscV = iota
	SUB
	MUL
	LW
	SW
)

const (
	Ready Status = iota
	Busy
	resultAvailable
)


// ReservationStation represents a reservation station in Tomasulo algorithm
type Register struct {
	RegisterID string
	Value string
	Status Status
}

// Instruction represents an instruction in Tomasulo algorithm
type Instruction struct {
	Opcode riscV // Operation (e.g., ADD, SUB, MUL)
	DestinationReg    string // Destination register
	SourceReg1    string // Source register 1
	SourceReg2    string // Source register 2
	Cycles int    // Execution cycles
} 

// ReservationStation represents a reservation station entry
type ReservationStation struct {
	Op          riscV // Operation type: add, subtract, multiply, etc.
	Rd string    // Destination register
	Vj, Vk      string    // Values of operands Vj and Vk
	Rj, Rk      string    // Reservation stations producing Vj and Vk (0 if ready)
    Busy        bool
    RemainingCycles int // Remaining execution cycles
    // Additional fields as needed
}