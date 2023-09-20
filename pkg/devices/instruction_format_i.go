package devices

import (
	"errors"

	"github.com/jzzheng22/riscvsim/pkg/instructions"
)

func (c *Cpu) DecodeIInstruction(instruction *instructions.Instruction) error {
	opcode, rd, rs1, funct3, imm, err := instruction.GetIFields()
	if err != nil {
		return err
	}

	regFile := c.GetRegFile()
	val1, err := regFile.GetReg(rs1)
	if err != nil {
		return err
	}

	result, err := iInstruction(c, opcode, funct3, val1, imm)
	if err != nil {
		return err
	}
	return regFile.SetReg(rd, result)
}

func iInstruction(c *Cpu, opcode, funct3 int32, val1 uint32, imm int32) (uint32, error) {
	switch opcode {
	// JALR
	case 0b1100111:
		return jalr(c, val1, uint32(imm))
	// Load
	case 0b0000011:
		// TODO: Implement Load instructions
		return load()
	// Immediate
	case 0b0010011:
		return immediate(funct3, val1, imm)
	// TODO: Implement FENCE instructions
	// FENCE
	case 0b0001111:
		return 0, errors.New("FENCE instructions not implemented")
	// TODO: Implement SYSTEM instructions
	// SYSTEM
	case 0b1110011:
		return 0, errors.New("SYSTEM instructions not implemented")
	default:
		return 0, &instructions.ExceptionIllegalInstruction{}
	}
}

func jalr(c *Cpu, val1, imm uint32) (uint32, error) {
	targetAddr := (val1 + imm) & (^uint32(1))
	if targetAddr%4 != 0 {
		return 0, &instructions.ExceptionMisalignedInstructionFetch{}
	}
	result := c.GetPc() + 4
	c.SetPc(targetAddr)
	return result, nil
}

// TODO: Implement this function
func load() (uint32, error) {
	return 0, errors.New("load not implemented")
}

func immediate(funct3 int32, val1 uint32, imm int32) (uint32, error) {
	switch funct3 {
	// ADDI
	case 0b000:
		return val1 + uint32(imm), nil
	// SLLI
	case 0b001:
		shamt := imm & 0b11111
		return val1 << shamt, nil
	// SLTI
	case 0b010:
		if int32(val1) < imm {
			return 1, nil
		}
		return 0, nil
	// SLTIU
	case 0b011:
		if val1 < uint32(imm) {
			return 1, nil
		}
		return 0, nil
	// XORI
	case 0b100:
		return val1 ^ uint32(imm), nil
	// SRLI, SRAI
	case 0b101:
		shamt := imm & 0b11111
		// SRLI
		if (imm>>10)&0b1 == 0 {
			return val1 >> shamt, nil
		}
		// SRAI
		return uint32(int32(val1) >> shamt), nil
	// ORI
	case 0b110:
		return val1 | uint32(imm), nil
	// ANDI
	case 0b111:
		return val1 & uint32(imm), nil
	default:
		return 0, errors.New("Impossible case in immediate()")
	}
}
