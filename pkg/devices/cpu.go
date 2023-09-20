package devices

type Cpu struct {
	regFile *RegisterFile
	memory  *Memory
	pc      uint32
}

func NewCpu() *Cpu {
	return &Cpu{
		regFile: &RegisterFile{
			registers: [32]uint32{},
		},
		memory: NewMemory(),
		pc:     0,
	}
}

func (c *Cpu) GetPC() uint32 {
	return c.pc
}

func (c *Cpu) GetMemory() *Memory {
	return c.memory
}

func (c *Cpu) GetRegFile() *RegisterFile {
	return c.regFile
}
