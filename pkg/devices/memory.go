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
	if addr > memorySize {
		return errors.New("SetWord() tried to access invalid memory address")
	}

	// Aligned on 4-byte boundary
	if addr%4 == 0 {
		m.memory[addr/4] = value
	} else {
		lowerIndex := addr / 4
		upperIndex := lowerIndex + 1
		lowerByteShift := (addr % 4) * 8
		upperByteShift := 32 - lowerByteShift

		m.memory[lowerIndex] &= (value << lowerByteShift) | (0xFFFFFFFF >> upperByteShift)
		m.memory[upperIndex] &= (value >> upperByteShift) | (0xFFFFFFFF << lowerByteShift)
	}
	return nil
}
