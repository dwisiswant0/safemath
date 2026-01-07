package safemath

// Integer is a constraint that permits any integer type.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// isSigned returns true if T is a signed integer type.
func isSigned[T Integer]() bool {
	var zero T

	return ^zero < 0
}

// Add returns the sum of a and b, or an error if overflow occurs.
func Add[T Integer](a, b T) (T, error) {
	c := a + b
	if isSigned[T]() {
		// Signed overflow occurs if operands have the same sign and the result
		// has a different sign.
		if (a > 0 && b > 0 && c < 0) || (a < 0 && b < 0 && c > 0) {
			return 0, ErrOverflow
		}
	} else {
		// Unsigned overflow occurs if the result is less than one of the
		// operands. (c < a) implies overflow.
		if c < a {
			return 0, ErrOverflow
		}
	}

	return c, nil
}

// Sub returns the difference of a and b, or an error if overflow occurs.
func Sub[T Integer](a, b T) (T, error) {
	c := a - b
	if isSigned[T]() {
		// Signed overflow occurs if operands have different signs and the
		// result has a different sign than a.
		// e.g., pos - neg = neg (Overflow)
		//       neg - pos = pos (Underflow)
		if (a > 0 && b < 0 && c < 0) || (a < 0 && b > 0 && c > 0) {
			return 0, ErrOverflow
		}
	} else {
		// Unsigned overflow (underflow) occurs if a < b.
		if a < b {
			return 0, ErrOverflow
		}
	}

	return c, nil
}

// Mul returns the product of a and b, or an error if overflow occurs.
func Mul[T Integer](a, b T) (T, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	// Specific edge case for signed integers: MinInt * -1
	// The problem is that c/a != b check relies on division. If a is -1 and c
	// is MinInt, c/a panics on some archs (MinInt / -1). So we must check for
	// -1 * MinInt (or vice versa) before doing the division check.
	if isSigned[T]() {
		minOne := ^T(0)
		if (a == minOne && b == -b) || (b == minOne && a == -a) {
			return 0, ErrOverflow
		}
	}

	c := a * b

	// General overflow check (covers most cases).
	if c/a != b {
		return 0, ErrOverflow
	}

	return c, nil
}

// Div returns the quotient of a and b.
func Div[T Integer](a, b T) (T, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if isSigned[T]() {
		minOne := ^T(0)
		// Signed overflow: MinInt / -1
		if b == minOne && a != 0 && a == -a {
			return 0, ErrOverflow
		}
	}

	return a / b, nil
}

// Convert safely converts a value from one Integer type to another.
func Convert[To, From Integer](v From) (To, error) {
	to := To(v)

	if isSigned[From]() && !isSigned[To]() {
		// Signed -> Unsigned
		if v < 0 {
			return 0, ErrTruncation
		}
	}

	if !isSigned[From]() && isSigned[To]() {
		// Unsigned -> Signed
		if to < 0 {
			return 0, ErrTruncation
		}
	}

	if From(to) != v {
		return 0, ErrTruncation
	}

	return to, nil
}

// ConvertAny attempts to convert v (any integer type) into To.
//
// Returns ErrInvalidType when v is not an integer or cannot fit in To.
func ConvertAny[To Integer](v any) (To, error) {
	switch x := v.(type) {
	case int:
		return Convert[To](x)
	case int8:
		return Convert[To](x)
	case int16:
		return Convert[To](x)
	case int32:
		return Convert[To](x)
	case int64:
		return Convert[To](x)
	case uint:
		return Convert[To](x)
	case uint8:
		return Convert[To](x)
	case uint16:
		return Convert[To](x)
	case uint32:
		return Convert[To](x)
	case uint64:
		return Convert[To](x)
	case uintptr:
		return Convert[To](x)
	default:
		return 0, ErrInvalidType
	}
}

// MustAdd returns the sum of a and b on success. Panics on error.
func MustAdd[T Integer](a, b T) T {
	c, err := Add(a, b)
	if err != nil {
		panic(err)
	}

	return c
}

// MustSub returns the difference of a and b on success. Panics on error.
func MustSub[T Integer](a, b T) T {
	c, err := Sub(a, b)
	if err != nil {
		panic(err)
	}

	return c
}

// MustMul returns the product of a and b on success. Panics on error.
func MustMul[T Integer](a, b T) T {
	c, err := Mul(a, b)
	if err != nil {
		panic(err)
	}

	return c
}

// MustDiv returns the quotient of a and b on success. Panics on error.
func MustDiv[T Integer](a, b T) T {
	c, err := Div(a, b)
	if err != nil {
		panic(err)
	}

	return c
}

// MustConvert safely converts a value from one Integer type to another on success. Panics on error.
func MustConvert[To, From Integer](v From) To {
	c, err := Convert[To](v)
	if err != nil {
		panic(err)
	}

	return c
}

// MustConvertAny converts v (any integer type) into To or panics on error.
func MustConvertAny[To Integer](v any) To {
	c, err := ConvertAny[To](v)
	if err != nil {
		panic(err)
	}

	return c
}
