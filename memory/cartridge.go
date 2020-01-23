package memory

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Cartridge struct {
	rom []byte
}

func NewCartridge(filename string) *Cartridge {
	rom, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("unable to read rom file %w", err))
	}

	return &Cartridge{
		rom:  rom[16:], // ignore first 16 bytes of header info
	}
}

func (c *Cartridge) Size() uint16 {
	return 0x3fff // another magic value, since we load at 0xc000 we want to fill the remaining bytes without overflowing
}

func (c *Cartridge) Read(addr uint16) byte {
	return c.rom[addr]
}

func (c *Cartridge) Write(addr uint16, data byte) {
	log.Println("dont write to rom!")
}
