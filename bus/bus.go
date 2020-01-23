package bus

type Device interface {
	Size () uint16
	Read (uint16) byte
	Write (uint16, byte)
}

type busDevice struct {
	Device
	offset uint16
}

type Bus struct {
	devices []*busDevice
}

func NewBus() *Bus {
	return &Bus{devices:make([]*busDevice, 0)}
}

func (b *Bus) AttachDevice(offset uint16, device Device) {
	d := &busDevice{
		Device: device,
		offset: offset,
	}
	b.devices = append(b.devices, d)
}

func (b *Bus) Write (addr uint16, data byte) {
	for _, v := range b.devices {
		if addr >= v.offset && addr < v.offset + v.Size() {
			v.Write(addr - v.offset, data)
			return
		}
	}
}

func (b *Bus) Read (addr uint16) byte {
	for _, v := range b.devices {
		if addr >= v.offset && addr < v.offset + v.Size() {
			return v.Read(addr - v.offset)
		}
	}
	return 0xFF
}
