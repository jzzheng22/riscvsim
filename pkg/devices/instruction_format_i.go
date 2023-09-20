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
	err = regFile.SetReg(rd, result)
	if err != nil {
		return err
	}
	return nil
}

func iInstruction(c *Cpu, opcode, funct3 int32, val1 uint32, imm int32) (uint32, error) {
	switch opcode {
	// JALR
	case 0b1100111:
		return jalr(c, val1, uint32(imm))
	// Load
	case 0b0000011:
		return load()
	// Immediate
	case 0b0010011:
		return immediate()
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

func load() (uint32, error) {
	return 0, errors.New("load not implemented")
}

func immediate() (uint32, error) {
	return 0, errors.New("immediate not implemented")
}
