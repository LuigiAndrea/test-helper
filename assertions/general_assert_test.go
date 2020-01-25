package assertions

import (
	"errors"
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
)

type testEqual struct {
	expected, actual interface{}
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
			t.Errorf("Test %d - %s", i+1, err.Error())
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
			t.Errorf("Test %d - Expected Exception!", i+1)
		}
	}
}

func TestAssertException(t *testing.T) {
	tests := []testException{
		testException{expected: &ValueError{}, actual: &ValueError{}},
		testException{expected: &LengthError{Err: errors.New("Slices with different length")}, actual: &LengthError{}},
	}

	for i, test := range tests {
		if err := AssertException(test.expected, test.actual); err != nil {
			t.Errorf("Test %d - %#v", i+1, err)
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
			t.Errorf("Test %d - Expected Exception!", i+1)
		} else if err.Error() != m.ErrorMessage(test.expected, test.actual) {
			t.Errorf("Test %d - Expected a different error message!", i+1)
		} else {
			var ee *ExceptionError
			if _, ok := err.(*ExceptionError); !ok {
				t.Errorf("Test %d - %s", i+1, m.ErrorMessage(ee, err))
			}
		}
	}
}
