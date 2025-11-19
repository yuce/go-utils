package types

import (
	"cmp"
	"sort"
)

type Pair[First cmp.Ordered, Second any] struct {
	First  First
	Second Second
}

func CollectMapToPairSlice[Key cmp.Ordered, Value any](m map[Key]Value) []Pair[Key, Value] {
	if m == nil {
		return nil
	}
	ls := make([]Pair[Key, Value], 0, len(m))
	for k, v := range m {
		ls = append(ls, Pair[Key, Value]{First: k, Second: v})
	}
	return ls
}

func SortPairSlice[First cmp.Ordered, Second any](pairs []Pair[First, Second]) {
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].First < pairs[j].First
	})
}
