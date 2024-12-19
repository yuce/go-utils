package must_test

import (
	"errors"
	"github.com/yuce/go-utils/must"
	"testing"
)

var (
	ErrSome = errors.New("some error")
)

func TestMust_NoErr(t *testing.T) {
	must.NoErr(f0NoErr())
}

func TestMust_Err(t *testing.T) {
	wrapper := func() (ret any) {
		defer func() {
			ret = recover()
		}()
		must.NoErr(f0Err())
		return nil
	}
	v := wrapper()
	if ve, ok := v.(error); ok {
		if !errors.Is(ve, ErrSome) {
			t.Fatalf("expected ErrSome, got: %v", ve)
		}
	} else {
		t.Fatalf("expected an error, got %v", v)
	}
}

func TestMustValue_NoErr(t *testing.T) {
	v := must.Value(f1NoErr())
	if v != "foo" {
		t.Fatalf("expected foo, got: %v", v)
	}
}

func TestMustValue_Err(t *testing.T) {
	wrapper := func() (ret any) {
		defer func() {
			ret = recover()
		}()
		v := must.Value(f1Err())
		return v
	}
	v := wrapper()
	if ve, ok := v.(error); ok {
		if !errors.Is(ve, ErrSome) {
			t.Fatalf("expected ErrSome, got: %v", ve)
		}
	} else {
		t.Fatalf("expected an error, got %v", v)
	}
}

func TestMustOK_True(t *testing.T) {
	v := must.OK(fOKTrue())
	if v != "foo" {
		t.Fatalf("expected foo, got: %v", v)
	}
}

func TestMustOK_False(t *testing.T) {
	wrapper := func() (ret any) {
		defer func() {
			ret = recover()
		}()
		v := must.OK(fOKFalse())
		return v
	}
	v := wrapper()
	if _, ok := v.(error); !ok {
		t.Fatalf("expected an error, got %v", v)
	}
}

func f0NoErr() error {
	return nil
}

func f0Err() error {
	return ErrSome
}

func f1NoErr() (any, error) {
	return "foo", nil
}

func f1Err() (string, error) {
	return "", ErrSome
}

func fOKTrue() (string, bool) {
	return "foo", true
}

func fOKFalse() (string, bool) {
	return "foo", false
}
