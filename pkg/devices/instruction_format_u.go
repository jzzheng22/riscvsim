package devices

import "github.com/jzzheng22/riscvsim/pkg/instructions"

func (c *Cpu) DecodeUInstruction(instruction *instructions.Instruction) error {
	opcode, rd, imm, err := instruction.GetUFields()
	if err != nil {
		return err
	}

	regFile := c.GetRegFile()
	result, err := uInstruction(c, opcode, rd, uint32(imm))
	if err != nil {
		return err
	}
	return regFile.SetReg(rd, result)
}

func uInstruction(c *Cpu, opcode int32, rd int, imm uint32) (uint32, error) {
	switch opcode {
	// LUI
	case 0b0110111:
		return imm, nil
	// AUIPC
	case 0b0010111:
		return c.GetPc() + imm, nil
	default:
		return 0, &instructions.ExceptionIllegalInstruction{}
	}
}
