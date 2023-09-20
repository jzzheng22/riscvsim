package devices

import "github.com/jzzheng22/riscvsim/pkg/instructions"

// TODO: Consider if return-address prediction stack behaviour is needed
func (c *Cpu) DecodeJInstruction(instruction *instructions.Instruction) error {
	rd, imm, err := instruction.GetJFields()
	if err != nil {
		return err
	}
	targetAddr := uint32(int32(c.GetPc()) + imm)
	if targetAddr%4 != 0 {
		return &instructions.ExceptionMisalignedInstructionFetch{}
	}
	err = c.GetRegFile().SetReg(rd, c.GetPc()+4)
	if err != nil {
		return err
	}
	c.SetPc(targetAddr)
	return nil
}
