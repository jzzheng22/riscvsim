package instructions

import "fmt"

type ExceptionIllegalInstruction struct{}

func (e *ExceptionIllegalInstruction) Error() string {
	return "Exception Code 0x2: Illegal instruction"
}

type ErrorWrongFormat struct {
	msg string
}

func NewErrWrongFormat(f Format, field string) *ErrorWrongFormat {
	return &ErrorWrongFormat{
		msg: fmt.Sprintf("Error: %s does not have field %s", f.String(), field),
	}
}

func (e *ErrorWrongFormat) Error() string {
	return e.msg
}

type ErrorWrongFields struct {
	msg string
}

func (e *ErrorWrongFields) Error() string {
	return e.msg
}

func NewErrWrongFields(target string, actual Format) *ErrorWrongFields {
	return &ErrorWrongFields{
		msg: fmt.Sprintf("Error: Trying to get %s fields from %s", target, actual.String()),
	}
}

type ExceptionMisalignedInstructionFetch struct{}

func (e *ExceptionMisalignedInstructionFetch) Error() string {
	return "Exception Code 0x0: Instruction address misaligned"
}
