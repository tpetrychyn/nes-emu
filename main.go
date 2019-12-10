package main

import (
	"github.com/tpetrychyn/nes-emu/bus"
	"github.com/tpetrychyn/nes-emu/cpu"
	"github.com/tpetrychyn/nes-emu/memory"
	"log"
	"time"
)

func main() {
	b := bus.NewBus()

	cart := memory.NewCartridge("nestest.nes")
	b.AttachDevice(0xc000, cart) // 0xc000 is where nestest starts reading from

	c := &cpu.Cpu{
		PC: 0xc000,
	}

	for {
		in := b.Read(c.PC)

		log.Printf("PC:%4X In:%2X A:%2X X:%2X Y:%2X - SR: %2X", c.PC, in, c.A, c.X, c.Y, c.SR)

		switch in {
		case 0xFF:
			return
		case 0x4A: // LSR Accumulator - shift all bits right 1 pos
			c.SetStatus(1 << 0, (c.A&1) == 1)
			c.A >>= 1
			c.PC ++
		case 0xA2: // LDX Immediate - store immediate value in X
			c.X = b.Read(c.PC+1)
			c.PC += 2
		case 0x4C: // jmp absolute
			addr := uint16(b.Read(c.PC+2)) << 8 | uint16(b.Read(c.PC+1))
			c.PC = addr
		default:
			c.PC++
		}

		<- time.After(1 * time.Second)
	}
}

