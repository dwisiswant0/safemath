// Package safemath is a safe, generic, and robust integer math library for Go.
//
// It provides overflow-safe arithmetic (+, -, *, /) and safe type conversions,
// preventing common bugs like overflow, underflow, and silent truncation that
// standard Go operations might miss.
//
// Every arithmetic operation (+, -, *, /) is provided in two variants:
//   - Error-returning: returns (T, error) (e.g., [Add], [Sub], [Mul], [Div])
//   - Panicking: returns T and panics on failure (e.g., [MustAdd], [MustSub], etc.)
//
// For type conversions, [Convert] ensures that the value can be represented
// in the target type without data loss, handling both signed-to-unsigned and
// size-based truncation checks.
package safemath
