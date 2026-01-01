# safemath

[![tests](https://github.com/dwisiswant0/safemath/actions/workflows/tests.yaml/badge.svg?branch=master)](https://github.com/dwisiswant0/safemath/actions/workflows/tests.yaml)
[![Go Reference](https://pkg.go.dev/badge/go.dw1.io/safemath.svg)](https://pkg.go.dev/go.dw1.io/safemath)

`safemath` is a safe, generic, and robust integer math library for Go.

It provides overflow-safe arithmetic (`+`, `-`, `*`, `/`) and safe type conversions, preventing common bugs like overflow, underflow, and silent truncation that standard Go operations might miss.

## Features

* **Comprehensive generics**: works with all standard integer types.
* **Checked arithmetic**: [`Add`](https://pkg.go.dev/go.dw1.io/safemath#Add), [`Sub`](https://pkg.go.dev/go.dw1.io/safemath#Sub), [`Mul`](https://pkg.go.dev/go.dw1.io/safemath#Mul), [`Div`](https://pkg.go.dev/go.dw1.io/safemath#Div) functions return an error instead of allowing silent, dangerous wrapping.
* **Safe conversions**: [`Convert[To, From](v)`](https://pkg.go.dev/go.dw1.io/safemath#Convert) makes sure no data is lost during type conversion (e.g., checking bounds when casting larger types to smaller ones or signed to unsigned).
* **Panic APIs**: [`Must*`](https://pkg.go.dev/go.dw1.io/safemath#MustAdd) variants are available for situations where panicking on failure is preferred.
* **Adversarial safety**: robustly handles dangerous edge cases like $$MinInt / -1$$ and avoids hardware exceptions.

## Usage

### Installation

```bash
go get go.dw1.io/safemath
```

### Basic Arithmetic

```go
package main

import (
    "fmt"

    "go.dw1.io/safemath"
)

func main() {
    a, b := 100, 200

    // safe addition
    sum, err := safemath.Add(a, b)
    if err != nil {
        fmt.Println("Overflow:", err)
        return
    }
    fmt.Println("Sum:", sum)

    // safe multiplication
    prod, err := safemath.Mul(a, b)
    if err != nil {
        fmt.Println("Overflow:", err)
        return
    }

    fmt.Println("Product:", prod)
}
```

### Safe Conversion

```go
package main

import (
    "fmt"

    "go.dw1.io/safemath"
)

func main() {
    var x int = 1000

    // converting large int to byte (should fail)
    val, err := safemath.Convert[byte](x)
    if err != nil {
        fmt.Printf("Conversion failed: %v\n", err) // Output: integer type truncation
    } else {
        fmt.Println("Converted:", val)
    }
}
```

### Panic-on-Error (Must APIs)

```go
package main

import (
    "go.dw1.io/safemath"
)

func main() {
    // will panic if overflow occurs
    sum := safemath.MustAdd(1, 2)
    println(sum)
}
```

## Acknowledgements

This project is heavily inspired by [trailofbits/go-panikint](https://github.com/trailofbits/go-panikint), a modified Go compiler that inserts automatic overflow checks at compile time.

`safemath` adopts a similar rigorous approach to correctness, throwing in:
* **Comprehensive edge cases**: tested against the same critical boundaries (e.g., $$MinInt / -1$$) to make sure it's robust.
* **Truncation safety**: [`Convert[To, From](v)`](https://pkg.go.dev/go.dw1.io/safemath#Convert) logic mirrors the strict checking philosophy to prevent silent data loss during type conversion.

While Go-Panikint enforces safety via a custom compiler toolchain, `safemath` provides these same rigorous protections as a standard, zero-dependency, and portable Go library.

## License

**safemath** is released with ♡ by [**@dwisiswant0**](https://github.com/dwisiswant0) under the Apache 2.0 license. See [LICENSE](/LICENSE).
