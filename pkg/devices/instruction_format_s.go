package devices

import (
	"errors"

	"github.com/jzzheng22/riscvsim/pkg/instructions"
)

// TODO: Implement S instructions
func (c *Cpu) DecodeSInstruction(instruction *instructions.Instruction) error {
	rs1, rs2, funct3, imm, err := instruction.GetSFields()
	if err != nil {
		return err
	}

	regFile := c.GetRegFile()
	val1, err := regFile.GetReg(rs1)
	if err != nil {
		return err
	}
	val2, err := regFile.GetReg(rs2)
	if err != nil {
		return err
	}

	byteAddr := uint32(int32(val1) + imm)

	switch funct3 {
	// SB
	case 0b000:
		return errors.New("SB not implemented")
	// SH
	case 0b001:
		return errors.New("SH not implemented")
	// SW
	case 010:
		return errors.New("SW not implemented")
	default:
		return &instructions.ExceptionIllegalInstruction{}
	}
}
