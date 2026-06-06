package iters_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/yuce/go-utils/assert"
	"github.com/yuce/go-utils/iters"
	"github.com/yuce/go-utils/recovers"
)

func TestChunk(t *testing.T) {
	const maxSize = 20
	for chunkSize := 1; chunkSize <= maxSize; chunkSize++ {
		itFunc := func() iter.Seq[int] {
			return iters.IntRange(0, maxSize, 1)
		}
		t.Run(fmt.Sprintf("chunk-size %d", chunkSize), func(t *testing.T) {
			chunks := slices.Collect(iters.Chunk(itFunc(), chunkSize))
			targetCnt := maxSize / chunkSize
			if targetCnt*chunkSize < maxSize {
				targetCnt++
			}
			assert.IntEqual(targetCnt, len(chunks))
			var collected []int
			for i, c := range chunks {
				t.Logf("chunk %d: %v", i, c)
				targetChunkSize := chunkSize
				if i == targetCnt-1 {
					targetChunkSize = maxSize - i*chunkSize
				}
				assert.IntEqual(targetChunkSize, len(c))
				collected = append(collected, c...)
			}
			t.Logf("chunk size: %d; collected: %v", chunkSize, collected)
			target := slices.Collect(itFunc())
			assert.SliceEqual(target, collected)
		})
	}
}

func TestChunk_ChunkSizeBiggerThanItems(t *testing.T) {
	it := iters.IntRange(0, 5, 1)
	chunks := slices.Collect(iters.Chunk(it, 10))
	target := [][]int{{0, 1, 2, 3, 4}}
	assert.IntEqual(1, len(target))
	assert.SliceEqual(target[0], chunks[0])
}

func TestIntRange(t *testing.T) {
	testCases := []struct {
		name   string
		it     iter.Seq[int]
		target []int
	}{
		{
			name:   "forward range",
			it:     iters.IntRange(3, 7, 1),
			target: []int{3, 4, 5, 6},
		},
		{
			name:   "backward range",
			it:     iters.IntRange(7, 3, -1),
			target: []int{7, 6, 5, 4},
		},
		{
			name:   "no range",
			it:     iters.IntRange(0, 0, 0),
			target: nil,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			sl := slices.Collect(tc.it)
			assert.SliceEqual(tc.target, sl)
		})
	}
}

func TestIntRange_Panics(t *testing.T) {
	testCases := []struct {
		name string
		it   iter.Seq[int]
	}{
		{
			name: "zero step with forward range",
			it:   iters.IntRange(3, 7, 0),
		},
		{
			name: "negative step with forward range",
			it:   iters.IntRange(3, 7, -1),
		},
		{
			name: "zero step with backward range",
			it:   iters.IntRange(7, 3, 0),
		},
		{
			name: "positive step with backward range",
			it:   iters.IntRange(7, 3, 1),
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := recovers.Err(func() {
				slices.Collect(tc.it)
			})
			assert.True(err != nil)
		})
	}
}
