package recovers_test

import (
	"fmt"
	"github.com/yuce/go-utils/recovers"
	"testing"
)

func TestRecover_NoErr(t *testing.T) {
	err := recovers.Err(func() {
		fmt.Println("OK")
	})
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
}

func TestRecover_Err(t *testing.T) {
	err := recovers.Err(func() {
		panics()
	})
	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}
}

func TestRecoverValue_NoErr(t *testing.T) {
	v, err := recovers.Value(func() string {
		return "OK"
	})
	if err != nil {
		t.Fatalf("expected nil, got: %v", err)
	}
	if v != "OK" {
		t.Fatalf("expected OK, got: %s", v)
	}
}

func TestRecoverValue_Err(t *testing.T) {
	_, err := recovers.Value(func() string {
		panics()
		return "FAIL"
	})
	if err == nil {
		t.Fatalf("expected error, got: %v", err)
	}
}

func panics() {
	panic("some panic")
}
