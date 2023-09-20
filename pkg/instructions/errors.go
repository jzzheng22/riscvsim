package instructions

import "fmt"

type ExceptionIllegalInstruction struct{}

func (e *ExceptionIllegalInstruction) Error() string {
	return "Exception Code 0x2: Illegal Instruction"
}

type ErrorWrongFormat struct {
	msg string
}

func NewErrWrongFormat(f Format, field string) *ErrorWrongFormat {
	return &ErrorWrongFormat{
		// TODO: Add stringer for Format and fix this call
		msg: fmt.Sprintf("Error: does not have field ", f, field),
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
		//TODO: add formatting to this
		msg: fmt.Sprintf("Error: Trying to get fields from ", target, actual),
	}
}
