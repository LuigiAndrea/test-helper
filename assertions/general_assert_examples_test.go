package assertions

import (
	"testing"

	"github.com/LuigiAndrea/test-helper/messages"
)

func ExampleAssertDeepEqual() {
	var t *testing.T

	type testEqual struct {
		expected, actual interface{}
	}
	tests := []testEqual{
		{expected: []int{1, 2}, actual: []int{1, 2}},
		{expected: [][]int{{1, 2}, {3, 3, 2}}, actual: [][]int{{1, 2}, {3, 3, 2}}},
		{expected: []string{"house", "apartment"}, actual: []string{"house", "apartment"}},
		{expected: 4.3, actual: 4.3},
	}

	for i, test := range tests {
		if err := AssertDeepEqual(test.expected, test.actual); err != nil {
			t.Errorf("Test %d - %s", i+1, err.Error())
		}
	}
}

func ExampleAssertNotDeepEqual() {
	var t *testing.T

	type testNotEqual struct {
		notExpected, actual interface{}
	}

	tests := []testNotEqual{
		{notExpected: []int{1, 3}, actual: []int{1, 2}},
		{notExpected: []string{"apartment"}, actual: []string{"house", "apartment"}},
		{notExpected: 4.3, actual: 5},
	}

	for i, test := range tests {
		if err := AssertNotDeepEqual(test.notExpected, test.actual); err != nil {
			t.Error(messages.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func ExampleAssertException() {
	var t *testing.T

	type testException struct {
		expected, actual error
	}

	tests := []testException{
		{expected: &ValueError{}, actual: &ValueError{}},
		{expected: &LengthError{}, actual: &LengthError{}},
	}

	for i, test := range tests {
		if err := AssertException(test.expected, test.actual); err != nil {
			t.Error(messages.ErrorMessageTestCount(i+1, err))
		}
	}
}

func ExampleAssertDeepException() {
	var t *testing.T

	type testException struct {
		expected, actual error
	}

	tests := []testException{
		{expected: &ValueError{X: 1, Y: 2, Pos: 12}, actual: &ValueError{X: 1, Y: 2, Pos: 12}},
		{expected: &LengthError{}, actual: &LengthError{}},
	}

	for i, test := range tests {
		if err := AssertDeepException(test.expected, test.actual); err != nil {
			t.Error(messages.ErrorMessageTestCount(i+1, err))
		}
	}
}

func ExampleAssertPanic() {
	var t *testing.T

	if err := AssertPanic(myFunc); err != nil {
		t.Error(err)
	}
}
