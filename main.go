package main

import (
	"github.com/tpetrychyn/nes-emu/cpu"
)

func main() {
	testProg := []byte{0x4A, 0xA2, 0x44, 0xFF}

	c := &cpu.Cpu{
		A: 0x1,
	}

	// TODO: read instructions from testProg
	// observe changes to the CPU state with the helpful print:
	// log.Printf("PC:%4X In:%2X A:%2X X:%2X Y:%2X - SR: %2X", c.PC, in, c.A, c.X, c.Y, c.SR)

	// Hints: 0x4A = LSR Accumulator - shift all bits right 1 pos
	//        0xA2 = LDX Immediate - store immediate value in X
	//        0xFF = _end - stop the program
}

