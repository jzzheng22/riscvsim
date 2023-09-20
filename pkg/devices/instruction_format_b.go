package devices

import "github.com/jzzheng22/riscvsim/pkg/instructions"

func (c *Cpu) DecodeBInstruction(instruction *instructions.Instruction) error {
	rs1, rs2, funct3, imm, err := instruction.GetBFields()
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

	takeBranch, err := takeBranch(funct3, val1, val2)
	if err != nil {
		return err
	}
	if takeBranch {
		c.SetPc(uint32(int32(c.GetPc()) + imm))
	}
	return nil
}

func takeBranch(funct3 int32, val1, val2 uint32) (bool, error) {
	switch funct3 {
	// BEQ
	case 0b000:
		return val1 == val2, nil
	// BNE
	case 0b001:
		return val1 != val2, nil
	// BLT
	case 0b100:
		return int32(val1) < int32(val2), nil
	// BGE
	case 0b101:
		return int32(val1) >= int32(val2), nil
	// BLTU
	case 0b110:
		return val1 < val2, nil
	// BGEU
	case 0b111:
		return val1 >= val2, nil
	default:
		return false, &instructions.ExceptionIllegalInstruction{}
	}
}
