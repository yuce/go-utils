package as_test

import (
	"reflect"
	"testing"

	"github.com/yuce/go-utils/as"
	"github.com/yuce/go-utils/assert"
)

type Stringable struct{}

func (v Stringable) String() string {
	return "x"
}

type otherStringType string

type otherType int

func TestString(t *testing.T) {
	text := "some text"
	otherText := otherStringType("another text")
	value := otherType(38)
	//
	testCases := []struct {
		name   string
		input  any
		output string
		ok     bool
	}{
		{
			name:   "string",
			input:  "some text",
			output: "some text",
			ok:     true,
		},
		{
			name:   "pointer to string",
			input:  &text,
			output: text,
			ok:     true,
		},
		{
			name:   "nil string",
			input:  (*string)(nil),
			output: "",
			ok:     true,
		},
		{
			name:   "other string type",
			input:  otherText,
			output: string(otherText),
			ok:     true,
		},
		{
			name:   "pointer to other string type",
			input:  &otherText,
			output: string(otherText),
			ok:     true,
		},
		{
			name:   "nil other string type",
			input:  (*otherStringType)(nil),
			output: "",
			ok:     true,
		},
		{
			name:   "stringer",
			input:  Stringable{},
			output: "x",
			ok:     true,
		},
		{
			name:   "nil stringer",
			input:  (*Stringable)(nil),
			output: "",
			ok:     true,
		},
		{
			name:   "non-stringer",
			input:  38,
			output: "",
			ok:     false,
		},
		{
			name:   "nil",
			input:  nil,
			output: "",
			ok:     false,
		},
		{
			name:   "pointer to non-string",
			input:  &value,
			output: "",
			ok:     false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, ok := as.String(tc.input)
			assert.True(reflect.DeepEqual(tc.output, s))
			assert.True(reflect.DeepEqual(tc.ok, ok))

		})
	}
}
