package measuring_test

import (
	"github.com/yuce/go-utils/measuring"
	"github.com/yuce/go-utils/must"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	const target = 10 * time.Millisecond
	dur := must.DurationNoError(measuring.Duration(func() error {
		time.Sleep(target)
		return nil
	}))
	if dur < target {
		t.Fatalf("duration too small: %v", dur)
	}
}

func TestDurationValue(t *testing.T) {
	const target = 10 * time.Millisecond
	dur, v := must.DurationValue(measuring.DurationValue(func() (int, error) {
		time.Sleep(target)
		return 42, nil
	}))
	if dur < target {
		t.Fatalf("duration too small: %v", dur)
	}
	if v != 42 {
		t.Fatalf("expected: 42, got: %v", v)
	}
}
