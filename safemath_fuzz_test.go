package safemath_test

import (
	"testing"

	"go.dw1.io/safemath"
)

// checkConsistency verifies that Must variant behaves exactly like the error-returning variant.
func checkConsistency[T any](t *testing.T, res T, err error, mustRes func() T) {
	t.Helper()
	if err != nil {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Operation failed with %v, but Must variant did not panic", err)
			}
		}()
		mustRes()
	} else {
		gotMust := mustRes()
		// Since T is generic and we can't always compare (e.g. if it's not comparable),
		// we use a simple comparison here as we only fuzz basic integer types.
		if any(gotMust) != any(res) {
			t.Errorf("Consistency mismatch: got %v, want %v", gotMust, res)
		}
	}
}

func FuzzAdd(f *testing.F) {
	f.Add(int64(1), int64(2))
	f.Fuzz(func(t *testing.T, a, b int64) {
		res, err := safemath.Add(a, b)
		checkConsistency(t, res, err, func() int64 { return safemath.MustAdd(a, b) })
	})
}

func FuzzSub(f *testing.F) {
	f.Add(int64(10), int64(5))
	f.Fuzz(func(t *testing.T, a, b int64) {
		res, err := safemath.Sub(a, b)
		checkConsistency(t, res, err, func() int64 { return safemath.MustSub(a, b) })
	})
}

func FuzzMul(f *testing.F) {
	f.Add(int64(10), int64(10))
	f.Fuzz(func(t *testing.T, a, b int64) {
		res, err := safemath.Mul(a, b)
		checkConsistency(t, res, err, func() int64 { return safemath.MustMul(a, b) })
	})
}

func FuzzDiv(f *testing.F) {
	f.Add(int64(100), int64(10))
	f.Fuzz(func(t *testing.T, a, b int64) {
		res, err := safemath.Div(a, b)
		checkConsistency(t, res, err, func() int64 { return safemath.MustDiv(a, b) })
	})
}

func FuzzArithmeticUint64(f *testing.F) {
	f.Add(uint64(10), uint64(5))
	f.Fuzz(func(t *testing.T, a, b uint64) {
		rAdd, eAdd := safemath.Add(a, b)
		checkConsistency(t, rAdd, eAdd, func() uint64 { return safemath.MustAdd(a, b) })

		rSub, eSub := safemath.Sub(a, b)
		checkConsistency(t, rSub, eSub, func() uint64 { return safemath.MustSub(a, b) })

		rMul, eMul := safemath.Mul(a, b)
		checkConsistency(t, rMul, eMul, func() uint64 { return safemath.MustMul(a, b) })

		if b != 0 {
			rDiv, eDiv := safemath.Div(a, b)
			checkConsistency(t, rDiv, eDiv, func() uint64 { return safemath.MustDiv(a, b) })
		}
	})
}

// FuzzConvertSignedToInt8 covers Signed -> Signed (smaller).
func FuzzConvertSignedToInt8(f *testing.F) {
	f.Add(int64(100))
	f.Fuzz(func(t *testing.T, a int64) {
		res, err := safemath.Convert[int8](a)
		checkConsistency(t, res, err, func() int8 { return safemath.MustConvert[int8](a) })
	})
}

// FuzzConvertSignedToUnsigned covers Signed -> Unsigned logic path.
func FuzzConvertSignedToUnsigned(f *testing.F) {
	f.Add(int64(-1))
	f.Fuzz(func(t *testing.T, a int64) {
		res, err := safemath.Convert[uint64](a)
		checkConsistency(t, res, err, func() uint64 { return safemath.MustConvert[uint64](a) })
	})
}

// FuzzConvertUnsignedToSigned covers Unsigned -> Signed logic path (overflow check).
func FuzzConvertUnsignedToSigned(f *testing.F) {
	f.Add(uint64(1) << 63)
	f.Fuzz(func(t *testing.T, a uint64) {
		res, err := safemath.Convert[int64](a)
		checkConsistency(t, res, err, func() int64 { return safemath.MustConvert[int64](a) })
	})
}

// FuzzConvertUnsignedToUnsignedSmall covers Unsigned -> Unsigned (smaller).
func FuzzConvertUnsignedToUnsignedSmall(f *testing.F) {
	f.Add(uint64(256))
	f.Fuzz(func(t *testing.T, a uint64) {
		res, err := safemath.Convert[uint8](a)
		checkConsistency(t, res, err, func() uint8 { return safemath.MustConvert[uint8](a) })
	})
}
