package strs_test

import (
	"fmt"
	"testing"

	"github.com/yuce/go-utils/assert"
	"github.com/yuce/go-utils/strs"
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

func TestSplitByComma(t *testing.T) {
	testCases := []struct {
		Name   string
		Text   string
		Result []string
	}{
		{
			Name:   "empty string",
			Text:   "",
			Result: nil,
		},
		{
			Name:   "single item",
			Text:   "foo",
			Result: []string{"foo"},
		},
		{
			Name:   "two items",
			Text:   "foo,bar",
			Result: []string{"foo", "bar"},
		},
		{
			Name:   "two items with space around comma",
			Text:   "foo, bar",
			Result: []string{"foo", "bar"},
		},
		{
			Name:   "three items",
			Text:   "foo, bar, quux",
			Result: []string{"foo", "bar", "quux"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			r := strs.SplitByComma(tc.Text)
			assert.SliceEqual(tc.Result, r)
		})
	}
}
