package as_test

import (
	"testing"

	"github.com/yuce/go-utils/as"
)

func BenchmarkString_String(b *testing.B) {
	var st any
	st = "x"
	for b.Loop() {
		as.String(st)
	}
}

func BenchmarkString_PtrString(b *testing.B) {
	st := "x"
	stp := &st
	for b.Loop() {
		as.String(stp)
	}
}

func BenchmarkString_StringLike(b *testing.B) {
	var st any
	st = otherStringType("x")
	for b.Loop() {
		as.String(st)
	}
}

func BenchmarkString_PtrStringLike(b *testing.B) {
	st := otherStringType("x")
	stp := &st
	for b.Loop() {
		as.String(stp)
	}
}

func BenchmarkString_Stringer(b *testing.B) {
	var st any
	st = Stringable{}
	for b.Loop() {
		as.String(st)
	}
}

func BenchmarkString_PtrStringer(b *testing.B) {
	st := Stringable{}
	stp := &st
	for b.Loop() {
		as.String(stp)
	}
}

func BenchmarkStringer(b *testing.B) {
	st := Stringable{}
	for b.Loop() {
		st.String()
	}
}
