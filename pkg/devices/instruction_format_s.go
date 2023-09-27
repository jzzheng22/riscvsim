package devices

import (
	"github.com/jzzheng22/riscvsim/pkg/instructions"
)

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
	memory := c.GetMemory()
	switch funct3 {
	// SB
	case 0b000:
		return memory.SetByte(byteAddr, uint8(val2))
	// SH
	case 0b001:
		return memory.SetHalf(byteAddr, uint16(val2))
	// SW
	case 010:
		return memory.SetWord(byteAddr, val2)
	default:
		return &instructions.ExceptionIllegalInstruction{}
	}
}
