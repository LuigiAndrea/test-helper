package assertions

import "testing"

func ExampleAssertDeepEqual() {
	var t *testing.T

	type testEqual struct {
		expected, actual interface{}
	}
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

func ExampleAssertException() {
	var t *testing.T

	type testEqual struct {
		expected, actual interface{}
	}

	tests := []testEqual{
		testEqual{expected: &ValueError{}, actual: &ValueError{}},
		testEqual{expected: &LengthError{}, actual: &LengthError{}},
	}

	for i, test := range tests {
		if err := AssertException(test.expected, test.actual); err != nil {
			t.Errorf("Test %d - %s", i+1, err.Error())
		}
	}
}
