package assertions

import (
	"math"
	"testing"

	"github.com/LuigiAndrea/test-helper/messages"
)

func ExampleAssertSlicesEqual_stringSlice() {
	var t *testing.T

	type testStringData struct {
		input1, input2 []string
	}

	tests := []testStringData{
		testStringData{input1: []string{"a", "b", "c", "d", "e"}, input2: []string{"a", "b", "c", "d", "e"}},
		testStringData{input1: []string{"c"}, input2: []string{"c"}},
		testStringData{input1: []string{""}, input2: []string{""}},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(StringSlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(messages.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func ExampleAssertSlicesEqual_dataSlice() {
	var t *testing.T

	type testData struct {
		input1, input2 []interface{}
	}

	tests := []testData{
		testData{input1: []interface{}{1, 32, 44, math.MaxInt32, 1}, input2: []interface{}{1, 32, 44, math.MaxInt32, 1}},
		testData{input1: []interface{}{true, false, "cat", 3}, input2: []interface{}{true, false, "cat", 3}},
		testData{input1: []interface{}{"AB", "c", "dog"}, input2: []interface{}{"AB", "c", "dog"}},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(DataSlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(messages.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}
