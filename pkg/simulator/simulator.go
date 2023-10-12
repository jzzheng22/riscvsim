package simulator

import (
	"errors"
	"os"

	"github.com/jzzheng22/riscvsim/pkg/devices"
	"github.com/jzzheng22/riscvsim/pkg/instructions"
)

func Simulate(binaryPath string) error {
	file, err := os.Open(binaryPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialise memory and register files
	cpu := devices.NewCpu()

	// TODO: Read binary file and load into memory

	// TODO: Execute instructions

	// TOOD: Fix this condition
	// Maybe use ECALL
	for cpu.GetPc() != 0 {
		err := simulateInstruction(cpu)
		if err != nil {
			return err
		}
		// TODO: Probably need to increment PC
	}

	return nil
}

func simulateInstruction(cpu *devices.Cpu) error {
	instructionWord, err := cpu.GetMemory().GetWord(cpu.GetPc())
	// TODO: handle error
	if err != nil {
		return err
	}

	instruction, exc := instructions.NewInstruction(int32(instructionWord))
	// TODO: handle exception
	if exc != nil {

	}
	switch instruction.GetFormat() {
	case instructions.FormatR:
		err := cpu.DecodeRInstruction(instruction)
		// TODO: Handle error
		if err != nil {
		}
	case instructions.FormatI:
		err := cpu.DecodeIInstruction(instruction)
		// TODO: Handle error
		if err != nil {
		}
	case instructions.FormatS:
		err := cpu.DecodeSInstruction(instruction)
		// TODO: Handle error
		if err != nil {
		}
		return errors.New("Not implemented")
	case instructions.FormatB:
		err := cpu.DecodeBInstruction(instruction)
		// TODO: Handle error
		if err != nil {
		}
		return errors.New("Not implemented")
	case instructions.FormatU:
		err := cpu.DecodeUInstruction(instruction)
		// TODO: Handle error
		if err != nil {
		}
	case instructions.FormatJ:
		err := cpu.DecodeJInstruction(instruction)
		// TODO: Handle error
		if err != nil {
		}
	}
	return nil
}
