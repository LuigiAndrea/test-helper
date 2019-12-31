package assertions

import (
	"testing"
)

type testDeepEqual struct {
	expected, actual interface{}
}

func TestAssertDeepEqual(t *testing.T) {
	tests := []testDeepEqual{
		testDeepEqual{expected: []int{1, 2}, actual: []int{1, 2}},
		testDeepEqual{expected: [][]int{{1, 2}, {3, 3, 2}}, actual: [][]int{{1, 2}, {3, 3, 2}}},
		testDeepEqual{expected: []string{"house", "apartment"}, actual: []string{"house", "apartment"}},
		testDeepEqual{expected: 4.3, actual: 4.3},
	}

	for i, test := range tests {
		if err := AssertDeepEqual(test.expected, test.actual); err != nil {
			t.Errorf("%d - %s", i+1, err.Error())
		}
	}
}

func TestAssertDeepEqualDifferentObjects(t *testing.T) {
	tests := []testDeepEqual{
		testDeepEqual{expected: []int{1}, actual: []int{1, 2}},
		testDeepEqual{expected: [][]int{{1, 2}, {3, 3, 2}}, actual: [][]int{{1, 2}, {3, 1, 2}}},
		testDeepEqual{expected: []int{1, 2}, actual: []string{"1", "2"}},
		testDeepEqual{expected: 4.3, actual: 4.31},
	}

	for i, test := range tests {
		if err := AssertDeepEqual(test.expected, test.actual); err == nil {
			t.Errorf("%d - Expected Exception!", i+1)
		} else {
			t.Log(err.Error())
		}
	}
}
