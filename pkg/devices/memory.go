package devices

import "errors"

type Memory struct {
	memory []uint32
}

const memorySize = 0x10000

func NewMemory() *Memory {
	return &Memory{
		memory: make([]uint32, memorySize/4),
	}
}

func (m *Memory) SetWord(addr, value uint32) error {
	if addr > 0x10000 {
		return errors.New("SetWord() tried to access invalid memory address")
	}
	m.memory[addr] = value
	return nil
}
