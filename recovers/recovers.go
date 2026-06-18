package recovers

import (
	"errors"
	"fmt"
)

// Err catches a panic and returns it as an error.
func Err(f func()) (err error) {
	defer func() {
		v := recover()
		if v != nil {
			switch vt := v.(type) {
			case error:
				err = vt
			case string:
				err = errors.New(vt)
			case fmt.Stringer:
				err = errors.New(vt.String())
			default:
				err = errors.New(fmt.Sprint(vt))
			}
		}
	}()
	f()
	return nil
}

// Value catches a panic and returns it as an error.
// If there are no panics, it returns the value.
func Value[T any](f func() T) (ret T, err error) {
	defer func() {
		v := recover()
		if v != nil {
			switch vt := v.(type) {
			case error:
				err = vt
			case string:
				err = errors.New(vt)
			case fmt.Stringer:
				err = errors.New(vt.String())
			default:
				err = errors.New(fmt.Sprint(vt))
			}
		}
	}()
	ret = f()
	return ret, nil
}
