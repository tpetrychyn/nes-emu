package cpu

type Cpu struct {
	A uint8 // accumulator register
	X uint8 // X register
	Y uint8 // Y register

	PC uint16 // program counter
	SP byte   // stack pointer
	SR byte   // status register
}

func (c *Cpu) SetStatus(bit uint8, state bool) {
	if state {
		c.SR |= bit // sets bit
	} else {
		c.SR &^= bit // clears bit
	}
}
