package main

import (
	"github.com/tpetrychyn/nes-emu/cpu"
	"log"
)

func main() {
	testProg := []byte{0x4A, 0xA2, 0x44, 0x99}

	c := &cpu.Cpu{
		A: 0x1,
	}

	for {
		in := testProg[c.PC]

		log.Printf("PC:%4X In:%2X A:%2X X:%2X Y:%2X - SR: %2X", c.PC, in, c.A, c.X, c.Y, c.SR)

		switch in {
		case 0x99:
			return
		case 0x4A: // LSR Accumulator - shift all bits right 1 pos
			c.SetStatus(1 << 0, (c.A&1) == 1)
			c.A >>= 1
			c.PC ++
		case 0xA2: // LDX Immediate - store immediate value in X
			c.X = testProg[c.PC+1]
			c.PC += 2
		default:
			c.PC++
		}
	}
}

