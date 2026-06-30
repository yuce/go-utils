package iters

import "iter"

// Chunk splits the given range into chunkSize items.
// Panics if chunkSize is not positive.
func Chunk[T any](it iter.Seq[T], chunkSize int) iter.Seq[[]T] {
	if chunkSize <= 0 {
		panic("iters.Chunk: chunkSize must be positive")
	}
	return func(yield func([]T) bool) {
		chunk := make([]T, 0, chunkSize)
		for item := range it {
			chunk = append(chunk, item)
			if len(chunk) == chunkSize {
				if !yield(chunk) {
					return
				}
				chunk = make([]T, 0, chunkSize)
			}
		}
		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}

// IntRange returns an iterator that counts from start to stop (excluded) by the given step.
// step must be positive if start < stop.
// step must be negative if start > stop.
func IntRange[T ~int](start, stop, step T) iter.Seq[T] {
	return func(yield func(T) bool) {
		if start < stop {
			if step <= 0 {
				panic("iters.IntRange: step must be positive if start < stop")
			}
			for start < stop {
				if !yield(start) {
					break
				}
				start += step
			}
		} else if start > stop {
			if step >= 0 {
				panic("iters.IntRange: step must be negative if start > stop")
			}
			for start > stop {
				if !yield(start) {
					break
				}
				start += step
			}
		}
	}
}

// FromChan creates an iterator from a channel
func FromChan[T any](ch <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range ch {
			if !yield(item) {
				return
			}
		}
	}
}

// Collect collects the items in the given iterator.
// It returns the partial list of items and the error if it encounters an error.
func Collect[T any](it iter.Seq[ErrItem[T]]) ([]T, error) {
	var r []T
	for ei := range it {
		if ei.Error != nil {
			return nil, ei.Error
		}
		r = append(r, ei.Item)
	}
	return r, nil
}

// ErrItem contains a value or an error
type ErrItem[T any] struct {
	Item  T
	Error error
}
