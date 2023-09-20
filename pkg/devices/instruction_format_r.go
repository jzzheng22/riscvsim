package devices

import (
	"github.com/jzzheng22/riscvsim/pkg/exceptions"
	"github.com/jzzheng22/riscvsim/pkg/instructions"
)

func (c *Cpu) DecodeRInstruction(instruction *instructions.Instruction) (*exceptions.Exception, error) {
	funct7, err := instruction.GetFunct7()
	if err != nil {
		return nil, err
	}

	funct3, err := instruction.GetFunct3()
	if err != nil {
		return nil, err
	}

	rd, err := instruction.GetRd()
	if err != nil {
		return nil, err
	}

	rs1, err := instruction.GetRs1()
	if err != nil {
		return nil, err
	}

	rs2, err := instruction.GetRs2()
	if err != nil {
		return nil, err
	}

	regFile := c.GetRegFile()
	val1, err := regFile.GetReg(rs1)
	// TODO: Handle this error
	if err != nil {

	}
	val2, err := regFile.GetReg(rs2)
	// TODO: Handle this error
	if err != nil {

	}

	result, exc := rInstruction(funct7, funct3, val1, val2)
	if exc != nil {
		return exc, nil
	}
	err = regFile.SetReg(rd, result)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func rInstruction(funct7, funct3 int32, val1, val2 uint32) (uint32, *exceptions.Exception) {
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
			return 0, &exceptions.Exception{
				ExceptionCode: exceptions.ExceptionIllegalInstruction,
			}
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
			return 0, &exceptions.Exception{
				ExceptionCode: exceptions.ExceptionIllegalInstruction,
			}
		}
	default:
		return 0, &exceptions.Exception{
			ExceptionCode: exceptions.ExceptionIllegalInstruction,
		}
	}
	// return nil, nil
}
