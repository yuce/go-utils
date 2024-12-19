package check

import "fmt"

// Recover catches a panic and returns it as an error.
func Recover(f func()) (err error) {
	defer func() {
		v := recover()
		if v != nil {
			if e, ok := v.(error); ok {
				err = fmt.Errorf("recover: %w", e)
				return
			}
			err = fmt.Errorf("recover: %s", v)
		}
	}()
	f()
	return err
}

// RecoverValue catches a panic and returns it as an error.
// If there are no panics, it returns the value.
func RecoverValue[T any](f func() T) (ret T, err error) {
	defer func() {
		v := recover()
		if v != nil {
			if e, ok := v.(error); ok {
				err = fmt.Errorf("recover: %w", e)
				return
			}
			err = fmt.Errorf("recover: %s", v)
		}
	}()
	ret = f()
	return ret, err
}
