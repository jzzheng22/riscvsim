package devices

import "fmt"

type RegisterFile struct {
	registers [32]uint32
}

type NonExistentRegError struct {
	regNum int
}

// zero		x0		Constant value 0
// ra		x1		Return address
// sp		x2		Stack pointer
// gp		x3		Global pointer
// tp		x4		Thread pointer
// t0-2		x5-7	Temporaries
// s0/fp	x8		Saved register / Frame pointer
// s1		x9		Saved register
// a0-1		x10-11	Function arguments / return values
// a2-7		x12-17	Function arguments
// s2-11	x18-27	Saved registers
// t3-6		x28-31	Temporaries

type Register int

const (
	RegisterZero Register = iota
	RegisterRA
	RegisterSP
	RegisterGP
	RegisterTP
	RegisterT0
	RegisterT1
	RegisterT2
	RegisterS0
	RegisterS1
	RegisterA0
	RegisterA1
	RegisterA2
	RegisterA3
	RegisterA4
	RegisterA5
	RegisterA6
	RegisterA7
	Registers2
	RegisterS3
	RegisterS4
	RegisterS5
	RegisterS6
	RegisterS7
	RegisterS8
	RegisterS9
	RegisterS10
	RegisterS11
	RegisterT3
	RegisterT4
	RegisterT5
	RegisterT6
)

const (
	fp Register = RegisterS0
)

func (r *NonExistentRegError) Error() string {
	return fmt.Sprintf("tried to set register x%d", r.regNum)
}

func (r *RegisterFile) SetReg(regNum int, value uint32) error {
	if regNum < 0 || regNum > 31 {
		return &NonExistentRegError{regNum}
	}
	//x0 hardwired to 0
	if regNum == 0 {
		return nil
	}
	r.registers[regNum] = value
	return nil
}

func (r *RegisterFile) GetReg(regNum int) (uint32, error) {
	if regNum < 0 || regNum > 31 {
		return 0, &NonExistentRegError{regNum}
	}
	return r.registers[regNum], nil
}
