package safemath_test

import (
	"math"
	"testing"

	"go.dw1.io/safemath"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name      string
		fn        func() error
		wantError error
	}{
		{
			name: "int8 overflow",
			fn: func() error {
				_, err := safemath.Add[int8](math.MaxInt8, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int8 underflow",
			fn: func() error {
				_, err := safemath.Add[int8](math.MinInt8, -1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "uint8 overflow",
			fn: func() error {
				_, err := safemath.Add[uint8](math.MaxUint8, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int64 overflow",
			fn: func() error {
				_, err := safemath.Add[int64](math.MaxInt64, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int64 underflow",
			fn: func() error {
				_, err := safemath.Add[int64](math.MinInt64, -1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "uint64 overflow",
			fn: func() error {
				_, err := safemath.Add[uint64](math.MaxUint64, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int8 ok",
			fn: func() error {
				got, err := safemath.Add[int8](1, 2)
				if err != nil {
					return err
				}
				if got != 3 {
					t.Errorf("want 3, got %d", got)
				}
				return nil
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err != tt.wantError {
				t.Errorf("got error %v, want %v", err, tt.wantError)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name      string
		fn        func() error
		wantError error
	}{
		{
			name: "int8 overflow (min - 1)",
			fn: func() error {
				_, err := safemath.Sub[int8](math.MinInt8, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int8 underflow (max - (-1))",
			fn: func() error {
				_, err := safemath.Sub[int8](math.MaxInt8, -1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "uint8 overflow (0 - 1)",
			fn: func() error {
				_, err := safemath.Sub[uint8](0, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int64 overflow (min - 1)",
			fn: func() error {
				_, err := safemath.Sub[int64](math.MinInt64, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "uint64 overflow (0 - 1)",
			fn: func() error {
				_, err := safemath.Sub[uint64](0, 1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err != tt.wantError {
				t.Errorf("got error %v, want %v", err, tt.wantError)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name      string
		fn        func() error
		wantError error
	}{
		{
			name: "int8 overflow",
			fn: func() error {
				_, err := safemath.Mul[int8](math.MaxInt8, 2)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int8 min * -1",
			fn: func() error {
				_, err := safemath.Mul[int8](math.MinInt8, -1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int8 -1 * min",
			fn: func() error {
				_, err := safemath.Mul[int8](-1, math.MinInt8)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int64 min * -1",
			fn: func() error {
				_, err := safemath.Mul[int64](math.MinInt64, -1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int64 -1 * min",
			fn: func() error {
				_, err := safemath.Mul[int64](-1, math.MinInt64)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int64 overflow",
			fn: func() error {
				_, err := safemath.Mul[int64](math.MaxInt64, 2)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int8 ok",
			fn: func() error {
				v, err := safemath.Mul[int8](10, 10)
				if err != nil {
					return err
				}
				if v != 100 {
					t.Errorf("want 100, got %d", v)
				}
				return nil
			},
			wantError: nil,
		},
		{
			name: "mul by zero",
			fn: func() error {
				v, err := safemath.Mul(10, 0)
				if err != nil {
					return err
				}
				if v != 0 {
					t.Errorf("want 0, got %d", v)
				}
				return nil
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err != tt.wantError {
				t.Errorf("got error %v, want %v", err, tt.wantError)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name      string
		fn        func() error
		wantError error
	}{
		{
			name: "divide by zero",
			fn: func() error {
				_, err := safemath.Div(1, 0)
				return err
			},
			wantError: safemath.ErrDivisionByZero,
		},
		{
			name: "int8 min / -1",
			fn: func() error {
				_, err := safemath.Div[int8](math.MinInt8, -1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
		{
			name: "int64 min / -1",
			fn: func() error {
				_, err := safemath.Div[int64](math.MinInt64, -1)
				return err
			},
			wantError: safemath.ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err != tt.wantError {
				t.Errorf("got error %v, want %v", err, tt.wantError)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name      string
		fn        func() error
		wantError error
	}{
		{
			name: "int to uint (negative)",
			fn: func() error {
				_, err := safemath.Convert[uint](-1)
				return err
			},
			wantError: safemath.ErrTruncation,
		},
		{
			name: "uint to int (overflow)",
			fn: func() error {
				// math.MaxUint64 to int64 should fail
				// Assume 64-bit arch for test or use explicit sizes
				_, err := safemath.Convert[int8, uint8](math.MaxUint8) // 255 -> -1
				return err
			},
			wantError: safemath.ErrTruncation,
		},
		{
			name: "uint64 to int64 overflow",
			fn: func() error {
				// math.MaxInt64 + 1 as uint64 is 1<<63. MinInt64 as int64.
				// But we are converting to int64.
				// 1<<63 cannot be represented in int64 (Max is 1<<63 - 1).
				val := uint64(math.MaxInt64) + 1
				_, err := safemath.Convert[int64](val)
				return err
			},
			wantError: safemath.ErrTruncation, // Truncation because it doesn't fit
		},
		{
			name: "int64 to uint64 negative",
			fn: func() error {
				_, err := safemath.Convert[uint64, int64](-1)
				return err
			},
			wantError: safemath.ErrTruncation,
		},
		{
			name: "large to small",
			fn: func() error {
				_, err := safemath.Convert[uint8, uint16](256)
				return err
			},
			wantError: safemath.ErrTruncation,
		},
		{
			name: "ok",
			fn: func() error {
				v, err := safemath.Convert[uint](10)
				if err != nil {
					return err
				}
				if v != 10 {
					t.Errorf("want 10, got %d", v)
				}
				return nil
			},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err != tt.wantError {
				t.Errorf("got error %v, want %v", err, tt.wantError)
			}
		})
	}
}

func TestMulPanicSafety(t *testing.T) {
	// Ensure that multiplying -1 by MinInt doesn't panic.
	// This specifically validates the fix for the "Mul Division Hazard" where
	// c/a checks could trigger a SIGFPE if executed before the safety check.

	try := func(name string, fn func() error) {
		t.Helper()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("%s panicked: %v", name, r)
			}
		}()
		err := fn()
		if err != safemath.ErrOverflow {
			t.Errorf("%s: want ErrOverflow, got %v", name, err)
		}
	}

	try("int8", func() error { return selectError(safemath.Mul[int8](-1, math.MinInt8)) })
	try("int16", func() error { return selectError(safemath.Mul[int16](-1, math.MinInt16)) })
	try("int32", func() error { return selectError(safemath.Mul[int32](-1, math.MinInt32)) })
	try("int64", func() error { return selectError(safemath.Mul[int64](-1, math.MinInt64)) })
	try("int", func() error { return selectError(safemath.Mul(-1, math.MinInt)) })
}

func selectError[T any](_ T, err error) error {
	return err
}

func TestMust(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		// Success case
		if got := safemath.MustAdd(1, 2); got != 3 {
			t.Errorf("MustAdd(1, 2) = %d; want 3", got)
		}
		// Panic case
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustAdd did not panic on overflow")
			} else {
				t.Logf("Recovered from: %v", r)
			}
		}()
		safemath.MustAdd[int8](math.MaxInt8, 1)
	})

	t.Run("Sub", func(t *testing.T) {
		// Success case
		if got := safemath.MustSub(3, 1); got != 2 {
			t.Errorf("MustSub(3, 1) = %d; want 2", got)
		}
		// Panic case
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustSub did not panic on overflow")
			} else {
				t.Logf("Recovered from: %v", r)
			}
		}()
		safemath.MustSub[int8](math.MinInt8, 1)
	})

	t.Run("Mul", func(t *testing.T) {
		// Success case
		if got := safemath.MustMul(2, 3); got != 6 {
			t.Errorf("MustMul(2, 3) = %d; want 6", got)
		}
		// Panic case
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustMul did not panic on overflow")
			} else {
				t.Logf("Recovered from: %v", r)
			}
		}()
		safemath.MustMul[int8](math.MaxInt8, 2)
	})

	t.Run("Div", func(t *testing.T) {
		// Success case
		if got := safemath.MustDiv(6, 2); got != 3 {
			t.Errorf("MustDiv(6, 2) = %d; want 3", got)
		}
		// Panic case
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustDiv did not panic on zero division")
			} else {
				t.Logf("Recovered from: %v", r)
			}
		}()
		safemath.MustDiv(1, 0)
	})

	t.Run("Convert", func(t *testing.T) {
		// Success case
		if got := safemath.MustConvert[uint](10); got != 10 {
			t.Errorf("MustConvert(10) = %d; want 10", got)
		}
		// Panic case
		defer func() {
			if r := recover(); r == nil {
				t.Error("MustConvert did not panic on truncation")
			} else {
				t.Logf("Recovered from: %v", r)
			}
		}()
		safemath.MustConvert[int8](128)
	})
}
