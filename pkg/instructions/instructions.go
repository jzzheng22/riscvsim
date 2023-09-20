package instructions

type Format int

const (
	FormatR Format = iota
	FormatI
	FormatS
	FormatB
	FormatU
	FormatJ
)

type Instruction struct {
	format Format
	opcode int32
	rd     int
	rs1    int
	rs2    int
	funct3 int32
	funct7 int32
	imm    int32
}

func NewInstruction(instruction int32) (*Instruction, error) {
	opcode := instruction & 0b1111111
	format, err := decodeFormat(opcode)
	if err != nil {
		return nil, err
	}
	return &Instruction{
		format: format,
		opcode: opcode,
		rd:     int((instruction >> 7) & 0b11111),
		rs1:    int((instruction >> 15) & 0b11111),
		rs2:    int((instruction >> 20) & 0b11111),
		funct3: (instruction >> 12) & 0b111,
		funct7: (instruction >> 25) & 0b1111111,
		imm:    decodeImmediate(instruction, format),
	}, nil
}

func decodeFormat(opcode int32) (Format, error) {
	switch opcode {
	// R-type instruction
	case 0110011:
		return FormatR, nil
	// JALR
	case 1100111:
		return FormatI, nil
	// Load
	case 0000011:
		return FormatI, nil
	// I-type immediate
	case 0010011:
		return FormatI, nil
	// TODO: Consider if this should be separate format
	// FENCE
	case 0001111:
		return FormatI, nil
	// TODO: Consider if this should be separate format
	// System instruction
	case 1110011:
		return FormatI, nil
	// S-type instruction
	case 0100011:
		return FormatS, nil
	// B-type instruction
	case 1100011:
		return FormatB, nil
	// LUI
	case 0110111:
		return FormatU, nil
	// AUIPC
	case 0010111:
		return FormatU, nil
	// JAL
	case 1101111:
		return FormatJ, nil
	// Opcode not recognised
	default:
		return 0, &ExceptionIllegalInstruction{}
	}
}

func decodeImmediate(instruction int32, format Format) int32 {
	switch format {
	case FormatI:
		return instruction >> 20
	case FormatS:
		imm11to5 := instruction >> 25
		imm4to0 := (instruction >> 7) & 0b11111
		return (imm11to5 << 5) | imm4to0
	case FormatB:
		imm12 := instruction >> 31
		imm10to5 := (instruction >> 25) & 0b111111
		imm4to1 := (instruction >> 8) & 0b1111
		imm11 := (instruction >> 7) & 0b1
		return (imm12 << 12) | (imm11 << 11) | (imm10to5 << 5) | (imm4to1 << 1)
	case FormatU:
		return instruction & 0b111111111111
	case FormatJ:
		imm20 := instruction >> 31
		imm10to1 := (instruction >> 21) & 0b1111111111
		imm11 := (instruction >> 20) & 0b1
		imm19to12 := (instruction >> 12) & 0b11111111
		return (imm20 << 20) | (imm19to12 << 12) | (imm11 << 11) | (imm10to1 << 1)
	default:
		return 0
	}
}

func (i *Instruction) GetFormat() Format {
	return i.format
}

func (i *Instruction) GetOpcode() int32 {
	return i.opcode
}
func (i *Instruction) GetRd() (int, error) {
	switch i.format {
	case FormatS, FormatB:
		return 0, NewErrWrongFormat(i.format, "rd")
	default:
		return i.rd, nil
	}
}

func (i *Instruction) GetRs1() (int, error) {
	switch i.format {
	case FormatU, FormatJ:
		return 0, NewErrWrongFormat(i.format, "rs1")
	default:
		return i.rs1, nil
	}
}

func (i *Instruction) GetRs2() (int, error) {
	switch i.format {
	case FormatI, FormatU, FormatJ:
		return 0, NewErrWrongFormat(i.format, "rs2")
	default:
		return i.rs2, nil
	}
}

func (i *Instruction) GetFunct3() (int32, error) {
	switch i.format {
	case FormatU, FormatJ:
		return 0, NewErrWrongFormat(i.format, "funct3")
	default:
		return i.funct3, nil
	}
}

func (i *Instruction) GetFunct7() (int32, error) {
	switch i.format {
	case FormatR:
		return i.funct7, nil
	default:
		return 0, NewErrWrongFormat(i.format, "funct7")
	}
}

func (i *Instruction) GetImm() (int32, error) {
	switch i.format {
	case FormatR:
		return 0, NewErrWrongFormat(i.format, "funct7")
	default:
		return i.imm, nil
	}
}
