package safemath_test

import (
	"testing"

	"go.dw1.io/safemath"
)

// Baseline benchmarks for native Go arithmetic to show overhead.

func BenchmarkAdd(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res = int64(i) + int64(i)
	}
	_ = res
}

func BenchmarkMul(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res = int64(i) * int64(i)
	}
	_ = res
}

func BenchmarkDiv(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res = int64(i) / 1
	}
	_ = res
}

// Safemath Arithmetic Benchmarks

func BenchmarkAddInt64(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Add(int64(i), int64(i))
	}
	_ = res
}

func BenchmarkAddUint64(b *testing.B) {
	var res uint64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Add(uint64(i), uint64(i))
	}
	_ = res
}

func BenchmarkMustAdd(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res = safemath.MustAdd(int64(i), int64(i))
	}
	_ = res
}

func BenchmarkSubInt64(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Sub(int64(i), int64(i))
	}
	_ = res
}

func BenchmarkSubUint64(b *testing.B) {
	var res uint64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Sub(uint64(i), uint64(i))
	}
	_ = res
}

func BenchmarkMustSub(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res = safemath.MustSub(int64(i), int64(i))
	}
	_ = res
}

func BenchmarkMulInt64(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Mul(int64(i), int64(i))
	}
	_ = res
}

func BenchmarkMulUint64(b *testing.B) {
	var res uint64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Mul(uint64(i), uint64(i))
	}
	_ = res
}

func BenchmarkMustMul(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res = safemath.MustMul(int64(i), int64(i))
	}
	_ = res
}

func BenchmarkDivInt64(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Div(int64(i), int64(1))
	}
	_ = res
}

func BenchmarkDivUint64(b *testing.B) {
	var res uint64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Div(uint64(i), uint64(1))
	}
	_ = res
}

func BenchmarkMustDiv(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res = safemath.MustDiv(int64(i), int64(1))
	}
	_ = res
}

// Safemath Conversion Benchmarks

func BenchmarkConvertSignedToInt8(b *testing.B) {
	var res int8
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Convert[int8](int64(i % 127))
	}
	_ = res
}

func BenchmarkConvertSignedToUint64(b *testing.B) {
	var res uint64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Convert[uint64](int64(i))
	}
	_ = res
}

func BenchmarkConvertUnsignedToSigned(b *testing.B) {
	var res int64
	for i := 0; b.Loop(); i++ {
		res, _ = safemath.Convert[int64](uint64(i))
	}
	_ = res
}

func BenchmarkMustConvert(b *testing.B) {
	var res int8
	for i := 0; b.Loop(); i++ {
		res = safemath.MustConvert[int8](int64(i % 127))
	}
	_ = res
}
