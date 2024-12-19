package must

import (
	"fmt"
)

// NoErr panics if err is not nil
func NoErr(err error) {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}
}

// Value panics if err is not nil.
// Otherwise, it returns v
func Value[T any](v T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}
	return v
}

// OK panics if ok is false.
// Otherwise, it returns v
func OK[T any](v T, ok bool) T {
	if !ok {
		panic(fmt.Errorf("must: not OK"))
	}
	return v
}
