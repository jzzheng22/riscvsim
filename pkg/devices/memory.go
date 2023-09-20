package devices

import "errors"

type Memory struct {
	memory []uint32
}

func NewMemory() *Memory {
	return &Memory{
		memory: make([]uint32, 0x10000),
	}
}

func (m *Memory) GetWord(addr uint32) (uint32, error) {
	if addr > 0x10000 {
		return 0, errors.New("GetWord() tried to access invalid memory address")
	}
	return m.memory[addr], nil
}

func (m *Memory) SetWord(addr, value uint32) error {
	if addr > 0x10000 {
		return errors.New("SetWord() tried to access invalid memory address")
	}
	m.memory[addr] = value
	return nil
}
