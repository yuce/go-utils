package assert

import (
	"fmt"
	"runtime"
)

func True(v bool) {
	if !v {
		loc := caller()
		panic(fmt.Errorf("assert: not true (%s)", loc))
	}
}

func False(v bool) {
	if v {
		loc := caller()
		panic(fmt.Errorf("assert: not false (%s)", loc))
	}
}

func IntEqual[T ~int | ~int8 | ~int16 | ~int32 | ~int64](v1, v2 T) {
	if v1 != v2 {
		loc := caller()
		panic(fmt.Errorf("assert: int %d != %d (%s)", v1, v2, loc))
	}
}

func SliceEqual[T comparable](v1, v2 []T) {
	if len(v1) != len(v2) {
		loc := caller()
		panic(fmt.Errorf("assert: slice length %d != %d (%s)", len(v1), len(v2), loc))
	}
	for i := range v1 {
		if v1[i] != v2[i] {
			loc := caller()
			panic(fmt.Errorf("assert: slice %v != %v (%s)", v1[i], v2[i], loc))
		}
	}
}

func StringEqual[T ~string](v1, v2 T) {
	if v1 != v2 {
		loc := caller()
		panic(fmt.Errorf("assert: string %s != %s (%s)", v1, v2, loc))
	}
}

func caller() string {
	_, fn, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%d", fn, line)
	}
	return "unknown"
}
