package safemath_test

import (
	"fmt"
	"math"

	"go.dw1.io/safemath"
)

func ExampleAdd() {
	// Normal addition
	sum, err := safemath.Add(10, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)

	// Overflow addition
	_, err = safemath.Add[int8](math.MaxInt8, 1)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 30
	// integer overflow/underflow
}

func ExampleMul() {
	// Normal multiplication
	prod, err := safemath.Mul(10, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prod)

	// Overflow multiplication
	_, err = safemath.Mul[int8](math.MaxInt8, 2)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 100
	// integer overflow/underflow
}

func ExampleSub() {
	// Normal subtraction
	diff, err := safemath.Sub(100, 30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(diff)

	// Underflow subtraction
	_, err = safemath.Sub[int8](math.MinInt8, 1)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 70
	// integer overflow/underflow
}

func ExampleDiv() {
	// Normal division
	quot, err := safemath.Div(100, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(quot)

	// Division by zero
	_, err = safemath.Div(100, 0)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 25
	// division by zero
}

func ExampleConvert() {
	// Safe conversion: int to int8 (success)
	v1, err := safemath.Convert[int8](100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v1)

	// Unsafe conversion: int to int8 (truncation)
	_, err = safemath.Convert[int8](200)
	if err != nil {
		fmt.Println(err)
	}

	// Unsafe conversion: negative signed to unsigned
	_, err = safemath.Convert[uint](-1)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 100
	// integer type truncation
	// integer type truncation
}

func ExampleConvertAny() {
	// Convert from interface{} holding an int64
	var v any = int64(42)
	res, err := safemath.ConvertAny[uint8](v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// Non-integer input is rejected
	_, err = safemath.ConvertAny[int]("nope")
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 42
	// invalid integer type
}

func ExampleMustAdd() {
	// MustAdd succeeds
	fmt.Println(safemath.MustAdd(100, 200))

	// MustAdd panics on overflow
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	safemath.MustAdd[int8](math.MaxInt8, 1)

	// Output:
	// 300
	// Recovered from: integer overflow/underflow
}

func ExampleMustConvert() {
	// MustConvert succeeds
	val := safemath.MustConvert[int16](int32(10000))
	fmt.Println(val)

	// MustConvert panics on truncation
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	safemath.MustConvert[int8](200)

	// Output:
	// 10000
	// Recovered from: integer type truncation
}

func ExampleMustConvertAny() {
	// MustConvertAny succeeds when the interface holds an integer
	var v any = uint32(255)
	fmt.Println(safemath.MustConvertAny[uint8](v))

	// MustConvertAny panics on invalid type
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	safemath.MustConvertAny[int]("bad")

	// Output:
	// 255
	// Recovered from: invalid integer type
}

func ExampleMustSub() {
	// MustSub succeeds
	fmt.Println(safemath.MustSub(100, 40))

	// MustSub panics on underflow
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	safemath.MustSub[uint8](0, 1)

	// Output:
	// 60
	// Recovered from: integer overflow/underflow
}

func ExampleMustMul() {
	// MustMul succeeds
	fmt.Println(safemath.MustMul(10, 10))

	// MustMul panics on overflow
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	safemath.MustMul[int8](100, 2)

	// Output:
	// 100
	// Recovered from: integer overflow/underflow
}

func ExampleMustDiv() {
	// MustDiv succeeds
	fmt.Println(safemath.MustDiv(100, 10))

	// MustDiv panics on division by zero
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	safemath.MustDiv(1, 0)

	// Output:
	// 10
	// Recovered from: division by zero
}
