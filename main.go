package main

import (
	"github.com/tpetrychyn/nes-emu/bus"
	"github.com/tpetrychyn/nes-emu/cpu"
	"github.com/tpetrychyn/nes-emu/memory"
	"time"
)

func main() {
	b := bus.NewBus()

	r := memory.NewRam(0x1fff)
	b.AttachDevice(0, r)

	cart := memory.NewCartridge("nestest.nes")
	b.AttachDevice(0xc000, cart) // 0xc000 is where nestest starts reading from

	c := cpu.NewCpu(b)

	c.PC = 0xc000

	for {
		res := c.Tick()

		if res == -1 {
			break
		}
		<- time.After(1 * time.Second)
	}
}

