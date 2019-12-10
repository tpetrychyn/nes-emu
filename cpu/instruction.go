package cpu

import (
	"fmt"
	"github.com/tpetrychyn/nes-emu/bus"
	"log"
)

type Instruction struct {
	*OpType

	op8 byte
	op16 byte
}

func ReadInstruction(pc uint16, bus *bus.Bus) (*Instruction, error) {
	i := bus.Read(pc)

	op, ok := optypes[i]
	if !ok {
		return nil, fmt.Errorf("bad instruction %2x", i)
	}

	in := &Instruction{
		OpType: op,
	}
	switch in.Bytes {
	case 1:
		break
	case 2:
		in.op8 = bus.Read(pc+1)
	case 3:
		in.op8 = bus.Read(pc+1)
		in.op16 = bus.Read(pc+2)
	default:
		log.Println("unhandled opcode byte length")
	}

	return in, nil
}