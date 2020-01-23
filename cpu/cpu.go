package cpu

import (
	"github.com/tpetrychyn/nes-emu/bus"
	"log"
)

const N = 1 << 7
const V = 1 << 6
const U = 1 << 5
const B = 1 << 4
const D = 1 << 3
const I = 1 << 2
const Z = 1 << 1
const C = 1 << 0

type Cpu struct {
	bus *bus.Bus
	A uint8 // accumulator register
	X uint8 // X register
	Y uint8 // Y register

	PC uint16 // program counter
	SP byte   // stack pointer
	SR byte   // status register

	in *Instruction
}

func (c *Cpu) setStatus(bit uint8, state bool) {
	if state {
		c.SR |= bit // sets bit
	} else {
		c.SR &^= bit // clears bit
	}
}

func (c *Cpu) setNZ(val uint8) {
	c.setStatus(N, (val >> 7) == 1)
	c.setStatus(Z, val == 0)
}

func NewCpu(bus *bus.Bus) *Cpu {
	return &Cpu{
		bus: bus,
	}
}

func (c *Cpu) getAddress() uint16 {
	switch c.in.addressing {
	case absolute:
		return uint16(c.in.op16) << 8 | uint16(c.in.op8)
	case absoluteX:
		return uint16(c.in.op16) << 8 | uint16(c.in.op8) + uint16(c.X)
	case absoluteY:
		return uint16(c.in.op16) << 8 | uint16(c.in.op8) + uint16(c.Y)

	case zeropage:
		return uint16(c.in.op8)
	case zeropageX:
		return uint16(c.in.op8) + uint16(c.X)
	case zeropageY:
		return uint16(c.in.op8) + uint16(c.Y)

	case indirect:
		pointer := uint16(c.in.op16) << 8 | uint16(c.in.op8)

		var addr uint16
		if c.in.op8 == 0x00FF {
			lo := uint16(c.bus.Read(pointer))
			hi := uint16(c.bus.Read(pointer & 0xFF00)) << 8

			addr = hi | lo
		} else {
			lo := uint16(c.bus.Read(pointer))
			hi := uint16(c.bus.Read(pointer + 1)) << 8

			addr = hi | lo
		}
		return addr

	case indirectX:
		pointer := uint16(c.in.op8) + uint16(c.X)
		lo := c.bus.Read(pointer)
		hi := c.bus.Read(pointer + 1)
		return uint16(hi) << 8 | uint16(lo)

	case indirectY:
		pointer := uint16(c.in.op8)
		lo := c.bus.Read(pointer)
		hi := c.bus.Read(pointer + 1)
		return (uint16(hi) << 8 | uint16(lo)) + uint16(c.Y)

	default:
		panic("bad address")
	}
}

func (c *Cpu) Tick() int {
	var err error
	c.in, err = ReadInstruction(c.PC, c.bus)
	if err != nil {
		log.Println(err)
		return -1
	}
	if c.in.Opcode == 0xFF { return -1 }
	log.Printf("PC:%4X In:%2X A:%2X X:%2X Y:%2X - SR: %2X", c.PC, c.in.Opcode, c.A, c.X, c.Y, c.SR)

	c.PC += uint16(c.in.Bytes)


	switch c.in.id {
	case lsr:
		c.LSR()
	case jmp:
		c.JMP()
	case ldx:
		c.LDX()
	case stx:
		c.STX()
	}

	c.in = nil
	return 0
}



func (c *Cpu) LSR() {
	if c.in.addressing == accumulator {
		c.setStatus(C, (c.A&1) == 1)
		c.A >>= 1
		c.setNZ(c.A)
		return
	}
}

func (c *Cpu) JMP() {
	c.PC = c.getAddress()
}

func (c *Cpu) LDX() {
	c.X = c.in.op8
	c.setNZ(c.X)
}

func (c *Cpu) STX() {
	addr := c.getAddress()
	c.bus.Write(addr, c.X)
}