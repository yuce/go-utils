package check

import (
	"fmt"
)

// Must panics if err is not nil
func Must(err error) {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}
}

// MustValue panics if err is not nil.
// Otherwise, it returns v
func MustValue[T any](v T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}
	return v
}

// MustOK panics if ok is false.
// Otherwise, it returns v
func MustOK[T any](v T, ok bool) T {
	if !ok {
		panic(fmt.Errorf("must: not OK"))
	}
	return v
}
