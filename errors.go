package safemath

import "errors"

var (
	ErrOverflow       = errors.New("integer overflow/underflow")
	ErrTruncation     = errors.New("integer type truncation")
	ErrInvalidType    = errors.New("invalid integer type")
	ErrDivisionByZero = errors.New("division by zero")
)
