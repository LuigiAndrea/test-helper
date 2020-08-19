package assertions

import (
	"errors"
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
)

type testEqual struct {
	expected, actual interface{}
}

type testNotEqual struct {
	notExpected, actual interface{}
}

type testException struct {
	expected, actual error
}

func TestAssertDeepEqual(t *testing.T) {
	tests := []testEqual{
		testEqual{expected: []int{1, 2}, actual: []int{1, 2}},
		testEqual{expected: [][]int{{1, 2}, {3, 3, 2}}, actual: [][]int{{1, 2}, {3, 3, 2}}},
		testEqual{expected: []string{"house", "apartment"}, actual: []string{"house", "apartment"}},
		testEqual{expected: 4.3, actual: 4.3},
	}

	for i, test := range tests {
		if err := AssertDeepEqual(test.expected, test.actual); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestAssertDeepEqualDifferentObjects(t *testing.T) {
	tests := []testEqual{
		testEqual{expected: []int{1}, actual: []int{1, 2}},
		testEqual{expected: [][]int{{1, 2}, {3, 3, 2}}, actual: [][]int{{1, 2}, {3, 1, 2}}},
		testEqual{expected: []int{1, 2}, actual: []string{"1", "2"}},
		testEqual{expected: 4.3, actual: 4.31},
	}

	for i, test := range tests {
		if err := AssertDeepEqual(test.expected, test.actual); err == nil {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected Exception!"))
		}
	}
}

func TestAssertNotDeepEqual(t *testing.T) {
	tests := []testNotEqual{
		testNotEqual{notExpected: []int{1, 3}, actual: []int{1, 2}},
		testNotEqual{notExpected: [][]int{{1, 2}, {3, 3, 2}}, actual: [][]int{{1, 2}, {3, 13, 2}}},
		testNotEqual{notExpected: []string{"apartment"}, actual: []string{"house", "apartment"}},
		testNotEqual{notExpected: 4.3, actual: 5},
	}

	for i, test := range tests {
		if err := AssertNotDeepEqual(test.notExpected, test.actual); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestAssertNotDeepEqualSameObjects(t *testing.T) {
	tests := []testNotEqual{
		testNotEqual{notExpected: []int{1, 3}, actual: []int{1, 3}},
		testNotEqual{notExpected: []string{"apartment"}, actual: []string{"apartment"}},
	}

	for i, test := range tests {
		if err := AssertNotDeepEqual(test.notExpected, test.actual); err == nil {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected Exception!"))
		}
	}
}

func TestAssertException(t *testing.T) {
	tests := []testException{
		testException{expected: &ValueError{}, actual: &ValueError{}},
		testException{expected: &LengthError{Err: "Slices with different length"}, actual: &LengthError{}},
	}

	for i, test := range tests {
		if err := AssertException(test.expected, test.actual); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestAssertDfferentException(t *testing.T) {
	tests := []testException{
		testException{expected: &ValueError{}, actual: nil},
		testException{expected: &ValueError{}, actual: &LengthError{}},
	}

	for i, test := range tests {
		if err := AssertException(test.expected, test.actual); err == nil {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected Exception!"))
		} else if err.Error() != m.ErrorMessage(test.expected, test.actual) {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected a different error message!"))
		} else {
			var e *ExceptionError
			if !errors.As(err, &e) {
				t.Error(m.ErrorMessageTestCount(i+1, m.ErrorMessage(e, err)))
			}
		}
	}
}

func TestAssertDeepException(t *testing.T) {
	tests := []testException{
		testException{expected: &ValueError{}, actual: &ValueError{}},
		testException{expected: &ValueError{X: 1, Y: 2, Pos: 12}, actual: &ValueError{X: 1, Y: 2, Pos: 12}},
		testException{expected: &LengthError{Err: "Slices with different length"},
			actual: &LengthError{Err: "Slices with different length"}},
		testException{expected: nil, actual: nil},
		testException{expected: &LengthError{}, actual: &LengthError{}},
	}

	for i, test := range tests {
		if err := AssertDeepException(test.expected, test.actual); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestAssertDifferentDeepException(t *testing.T) {
	tests := []testException{
		testException{expected: &ValueError{}, actual: &LengthError{}},
		testException{expected: &ValueError{X: 1, Y: 0, Pos: 12}, actual: &ValueError{X: 1, Y: 2, Pos: 12}},
		testException{expected: &LengthError{Err: "Slices"}, actual: &LengthError{Err: "SlicesB"}},
		testException{expected: nil, actual: &LengthError{}},
		testException{expected: &LengthError{}, actual: nil},
		testException{expected: &LengthError{}, actual: &ValueError{}},
	}

	for i, test := range tests {
		if err := AssertDeepException(test.expected, test.actual); err == nil {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected Exception!"))
		} else if err.Error() != m.ErrorMessage(test.expected, test.actual) {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected a different error message!"))
		} else {
			var e *ExceptionError
			if !errors.As(err, &e) {
				t.Error(m.ErrorMessageTestCount(i+1, m.ErrorMessage(e, err)))
			}
		}
	}
}

func TestAssertPanic(t *testing.T) {
	if err := AssertPanic(myFunc); err != nil {
		t.Error(err)
	}

	if err := AssertPanic(func() {
		myFuncParameter("myString")
	}); err != nil {
		t.Error(err)
	}
}

func myFunc() {
	panic("it is panicking")
}

func myFuncParameter(par string) {
	panic(par)
}
