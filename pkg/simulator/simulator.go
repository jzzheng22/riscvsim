package simulator

import (
	"errors"
	"os"

	"github.com/jzzheng22/riscvsim/pkg/devices"
	"github.com/jzzheng22/riscvsim/pkg/instructions"
)

func Simulate(binaryPath string) (int, error) {
	exitCode := 0

	file, err := os.Open(binaryPath)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	// Initialise memory and register files
	cpu := devices.NewCpu()

	// Read binary file and load into memory

	// Execute instructions

	// TOOD: Fix this condition
	// Maybe use ECALL
	for cpu.GetPc() != 0 {
		instructionWord, err := cpu.GetMemory().GetWord(cpu.GetPc())
		// TODO: handle error
		if err != nil {
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
			return 0, errors.New("Not implemented")
		case instructions.FormatB:
			return 0, errors.New("Not implemented")
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
	}
	// Set exit code

	return exitCode, nil
}
