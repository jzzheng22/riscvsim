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

func (m *Memory) SetHalf(addr uint32, value uint16) error {
	if addr > memorySize {
		return errors.New("SetHalf() tried to access invalid memory address")
	}
	switch addr % 4 {
	case 0:
		m.memory[addr/4] &= uint32(value) | 0xFFFF0000
	case 1:
		m.memory[addr/4] &= (uint32(value) << 8) | 0xFF0000FF
	case 2:
		m.memory[addr/4] &= (uint32(value) << 16) | 0xFFFF
		// return uint16((m.memory[addr/4] >> 16) & 0xFFFF), nil
	case 3:
		lowerHalf := uint32(value&0xFF) << 24
		upperHalf := uint32(value&0xFF00) >> 8
		m.memory[addr/4] &= lowerHalf | 0xFFFFFF
		m.memory[addr/4+1] &= upperHalf | 0xFFFFFF00
	default:
		return errors.New("SetHalf() entered illegal default case")
	}

	return nil
}


func (m *Memory) GetWord(addr uint32) (uint32, error) {
	if addr > memorySize {
		return 0, errors.New("GetWord() tried to access invalid memory address")
	}
	// Aligned on 4-byte boundary
	if addr%4 == 0 {
		return m.memory[addr/4], nil
	} else {
		lowerIndex := addr / 4
		upperIndex := lowerIndex + 1
		lowerByteShift := (addr % 4) * 8
		upperByteShift := 32 - lowerByteShift
		lowerWord := m.memory[lowerIndex] >> lowerByteShift
		upperWord := m.memory[upperIndex] << upperByteShift
		return upperWord | lowerWord, nil
	}
}

func (m *Memory) GetHalf(addr uint32) (uint16, error) {
	if addr > memorySize {
		return 0, errors.New("GetWord() tried to access invalid memory address")
	}
	switch addr % 4 {
	case 0:
		return uint16(m.memory[addr/4] & 0xFFFF), nil
	case 1:
		return uint16((m.memory[addr/4] >> 8) & 0xFFFF), nil
	case 2:
		return uint16((m.memory[addr/4] >> 16) & 0xFFFF), nil
	case 3:
		lowerWord := m.memory[addr/4] & 0xFF000000
		upperWord := m.memory[addr/4+1] & 0xFF
		return uint16(upperWord<<8 | lowerWord>>24), nil
	default:
		return 0, errors.New("GetHalf() entered illegal default case")
	}
}

func (m *Memory) GetByte(addr uint32) (uint8, error) {
	if addr > memorySize {
		return 0, errors.New("GetWord() tried to access invalid memory address")
	}
	word := m.memory[addr/4]
	shift := (addr % 4) * 8
	return uint8((word >> shift) & 0xFF), nil
}
