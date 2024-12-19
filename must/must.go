package must

import (
	"fmt"
	"time"
)

// NoError panics if err is not nil
func NoError(err error) {
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

// DurationNoError panics if err is false.
// Otherwise it returns duration.
func DurationNoError(duration time.Duration, err error) time.Duration {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}
	return duration
}

// DurationValue panics if err is false.
// Otherwise, it returns duration and v.
func DurationValue[T any](duration time.Duration, v T, err error) (time.Duration, T) {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}
	return duration, v
}
