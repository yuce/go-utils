package strs_test

import (
	"fmt"
	"github.com/yuce/go-utils/assert"
	"github.com/yuce/go-utils/strs"
	"testing"
)

func TestFirstNonEmpty(t *testing.T) {
	testCases := []struct {
		Strings []string
		Want    string
	}{
		{
			Strings: []string{},
			Want:    "",
		},
		{
			Strings: []string{""},
			Want:    "",
		},
		{
			Strings: []string{"a", "b", "c"},
			Want:    "a",
		},
		{
			Strings: []string{"", "b", "c"},
			Want:    "b",
		},
		{
			Strings: []string{"", "", "c"},
			Want:    "c",
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.Strings), func(t *testing.T) {
			o := strs.FirstNonEmpty(tc.Strings...)
			assert.StringEqual(tc.Want, o)
		})
	}
}
