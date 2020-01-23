package cpu

// addressing modes
const (
	_ = iota
	absolute
	absoluteX
	absoluteY
	accumulator
	immediate
	implied
	indirect
	indirectX
	indirectY
	relative
	zeropage
	zeropageX
	zeropageY
)

const (
	_ = iota
	_end
	lsr
	ldx
	jmp
	stx
)

type OpType struct {
	// Opcode is a byte representing an instruction and its addressing mode.
	Opcode byte

	// id is an internal identifier of the instruction type/mnemonic, e.g. ADC
	id uint8

	// addressing is an internal identifier for the addressing mode.
	addressing uint8

	// Bytes is the size of the instruction with its operand.
	// Opcodes with implicit/null operand are 1 byte.
	// Opcodes with immediate or zeropage operand are 2 bytes.
	// Opcodes with adddress operands are 3 bytes.
	Bytes uint8

	// Cycles is the number of times the system clock signal will rise and fall
	// before the instruction is complete.
	Cycles uint8
}

var optypes = map[uint8]*OpType{
	0xFF: {0xFF, _end, implied, 1, 1},
	0x4A: {0x4A, lsr, absolute, 1, 2},
	0xA2: {0xA2, ldx, immediate, 2, 2},
	0x4C: {0x4C, jmp, absolute, 3, 3},
	0x86: {0x86, stx, zeropage, 2, 2},
}