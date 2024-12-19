package measuring

import (
	"context"
	"time"
)

// Duration measures the time to execute f.
func Duration(f func() error) (time.Duration, error) {
	return DurationCtx(context.Background(), func(ctx context.Context) (err error) {
		return f()
	})
}

// DurationValue measures the time to execute f, and returns the result.
func DurationValue[T any](f func() (T, error)) (time.Duration, T, error) {
	return DurationValueCtx(context.Background(), func(ctx context.Context) (T, error) {
		return f()
	})
}

// DurationCtx measures the time to execute f.
func DurationCtx(ctx context.Context, f func(ctx context.Context) error) (time.Duration, error) {
	tic := time.Now()
	if err := f(ctx); err != nil {
		return 0, err
	}
	took := time.Since(tic)
	return took, nil
}

// DurationValueCtx measures the time to execute f, and returns the result.
func DurationValueCtx[T any](ctx context.Context, f func(ctx context.Context) (T, error)) (time.Duration, T, error) {
	tic := time.Now()
	v, err := f(ctx)
	if err != nil {
		return 0, v, err
	}
	took := time.Since(tic)
	return took, v, nil
}
