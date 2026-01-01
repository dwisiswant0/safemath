package safemath

import "errors"

var (
	ErrOverflow       = errors.New("integer overflow/underflow")
	ErrTruncation     = errors.New("integer type truncation")
	ErrDivisionByZero = errors.New("division by zero")
)
