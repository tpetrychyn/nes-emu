package memory

import "log"

type Ram struct {
	size uint16
	mem []byte
}

func NewRam(size uint16) *Ram {
	return &Ram{
		size: size,
		mem:  make([]byte, size),
	}
}

func (r *Ram) Size() uint16 {
	return r.size
}

func (r *Ram) Read(addr uint16) byte {
	if addr < r.size {
		return r.mem[addr & 0x07FF] // mirror ram every 0x07ff
	}

	return 0
}

func (r *Ram) Write(addr uint16, val byte) {
	log.Printf("writing ram value %2X at %2X", val, addr)
	r.mem[addr] = val
}
