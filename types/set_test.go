package types_test

import (
	"github.com/yuce/go-utils/assert"
	"github.com/yuce/go-utils/types"
	"slices"
	"testing"
)

func TestNewSet(t *testing.T) {
	s1 := types.NewSet[string]()
	assert.IntEqual(0, len(s1))
	assert.SliceEqual([]string{}, s1.Slice())
	s2 := types.NewSet(1, 2, 3)
	assert.IntEqual(3, len(s2))
	sl := s2.Slice()
	slices.Sort(sl)
	assert.SliceEqual(sl, s2.Slice())
}

func TestSet_Add(t *testing.T) {
	s := types.NewSet[string]()
	s.Add("foo")
	assert.IntEqual(1, len(s))
	assert.SliceEqual([]string{"foo"}, s.Slice())
	s.Add("bar")
	assert.IntEqual(2, len(s))
	sl := s.Slice()
	slices.Sort(sl)
	assert.SliceEqual([]string{"bar", "foo"}, sl)
	// add the same item
	s.Add("bar")
	assert.IntEqual(2, len(s))
	sl = s.Slice()
	slices.Sort(sl)
	assert.SliceEqual([]string{"bar", "foo"}, sl)
}

func TestSet_Remove(t *testing.T) {
	s := types.NewSet[string]()
	// remove non-existent item
	s.Remove("foo")
	assert.IntEqual(0, len(s))
	s.Add("foo")
}
