package devices

import (
	"github.com/jzzheng22/riscvsim/pkg/instructions"
)

func (c *Cpu) DecodeRInstruction(instruction *instructions.Instruction) error {
	rd, rs1, rs2, funct3, funct7, err := instruction.GetRFields()
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

	result, err := rInstruction(funct7, funct3, val1, val2)
	if err != nil {
		return err
	}
	return regFile.SetReg(rd, result)
}

func rInstruction(funct7, funct3 int32, val1, val2 uint32) (uint32, error) {
	switch funct7 {
	case 0b0000000:
		switch funct3 {
		// ADD
		case 0b000:
			return val1 + val2, nil
		// SLL
		case 0b001:
			return val1 << (val2 & 0b11111), nil
		// SLT
		case 0b010:
			if int32(val1) < int32(val2) {
				return 1, nil
			}
			return 0, nil
		// SLTU
		case 0b011:
			if val1 < val2 {
				return 1, nil
			}
			return 0, nil
		// XOR
		case 0b100:
			return val1 ^ val2, nil
		// SRL
		case 0b101:
			return val1 >> (val2 & 0b11111), nil
		// OR
		case 0b110:
			return val1 | val2, nil
		// AND
		case 0b111:
			return val1 & val2, nil
		default:
			return 0, &instructions.ExceptionIllegalInstruction{}
		}
	case 0b0100000:
		switch funct3 {
		// SUB
		case 0b000:
			return val1 - val2, nil
		// SRLA
		case 0b101:
			return uint32(int32(val1) >> (val2 & 0b11111)), nil
		default:
			return 0, &instructions.ExceptionIllegalInstruction{}
		}
	default:
		return 0, &instructions.ExceptionIllegalInstruction{}
	}
	// return nil, nil
}
