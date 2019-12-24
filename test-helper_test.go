package gotesthelper

import (
	"math"
	"testing"
)

type testStringData struct {
	input1, input2 []string
}

type testIntData struct {
	input1, input2 []int
}

type testFloatData struct {
	input1, input2 []float64
}

type testByteData struct {
	input1, input2 []byte
}

type testData struct {
	input1, input2 []interface{}
}

func TestHelperStringArray(t *testing.T) {
	tests := []testStringData{
		testStringData{input1: []string{"a", "b", "c", "d", "e"}, input2: []string{"a", "b", "c", "d", "e"}},
		testStringData{input1: []string{"c"}, input2: []string{"c"}},
		testStringData{input1: []string{""}, input2: []string{""}},
		testStringData{input1: []string{}, input2: []string{}},
		testStringData{input1: []string{"d", "e", "a"}, input2: []string{"d", "e", "a"}},
		testStringData{input1: []string{"T", "R", "E"}, input2: []string{"T", "R", "E"}},
		testStringData{input1: []string{"CC", "AA", "BB"}, input2: []string{"CC", "AA", "BB"}},
	}

	for i, test := range tests {
		if err := AssertArraysEqual(StringArrays{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Errorf("%d - %v", i+1, err.Error())
		}
	}
}

func TestHelperDifferentElementsStringArray(t *testing.T) {
	tests := []testStringData{
		testStringData{input1: []string{"-1", "1"}, input2: []string{"-1", "13"}},
		testStringData{input1: []string{"1", "5", "7", "8"}, input2: []string{"1", "5", "3", "8"}},
		testStringData{input1: []string{""}, input2: []string{" "}},
	}

	for i, test := range tests {
		err := AssertArraysEqual(StringArrays{Expected: test.input1, Actual: test.input2})

		if err == nil {
			t.Errorf("%d - Expected Exception! Array with different values", i+1)
		}

		_, isValueError := err.(*ValueError)

		if !isValueError {
			t.Errorf("%d - Expected a ValueError: %T", i+1, err)
		}
	}
}

func TestHelperIntArray(t *testing.T) {
	tests := []testIntData{
		testIntData{input1: []int{1, 32, 44322, math.MaxInt64, math.MinInt64}, input2: []int{1, 32, 44322, math.MaxInt64, math.MinInt64}},
		testIntData{input1: []int{133}, input2: []int{133}},
		testIntData{input1: []int{0}, input2: []int{0}},
		testIntData{input1: []int{}, input2: []int{}},
		testIntData{input1: []int{-3, 43, -0}, input2: []int{-3, 43, -0}},
	}

	for i, test := range tests {
		if err := AssertArraysEqual(IntArrays{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Errorf("%d - %v", i+1, err.Error())
		}
	}
}

func TestHelperDifferentElementsIntArray(t *testing.T) {
	tests := []testIntData{
		testIntData{input1: []int{-1, 3}, input2: []int{-1, 33}},
		testIntData{input1: []int{1, 5, 7, 8}, input2: []int{1, 5, 3, 8}},
	}

	for i, test := range tests {
		err := AssertArraysEqual(IntArrays{Expected: test.input1, Actual: test.input2})

		if err == nil {
			t.Errorf("%d - Expected Exception! Array with different values", i+1)
		}

		_, isValueError := err.(*ValueError)

		if !isValueError {
			t.Errorf("%d - Expected a ValueError: %T", i+1, err)
		}
	}
}

func TestHelperFloat64Array(t *testing.T) {

	tests := []testFloatData{
		testFloatData{input1: []float64{1, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(2)}, input2: []float64{1, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(2)}},
		testFloatData{input1: []float64{133}, input2: []float64{133}},
		testFloatData{input1: []float64{0.0}, input2: []float64{0.0}},
		testFloatData{input1: []float64{}, input2: []float64{}},
		testFloatData{input1: []float64{-3.0, 43.0, -0.0}, input2: []float64{-3.0, 43.0, -0.0}},
		testFloatData{input1: []float64{2.5, 3.3}, input2: []float64{2.5, 3.3}},
	}

	for i, test := range tests {
		if err := AssertArraysEqual(Float64Arrays{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Errorf("%d - %v", i+1, err.Error())
		}
	}
}

func TestHelperDifferentElementsFloatArray(t *testing.T) {

	tests := []testFloatData{
		testFloatData{input1: []float64{-1, 3}, input2: []float64{-1, 33}},
		testFloatData{input1: []float64{1, 5, 7, 8}, input2: []float64{1, 5, 3, 8}},
	}

	for i, test := range tests {
		err := AssertArraysEqual(Float64Arrays{Expected: test.input1, Actual: test.input2})

		if err == nil {
			t.Errorf("%d - Expected Exception! Array with different values", i+1)
		}

		_, isValueError := err.(*ValueError)

		if !isValueError {
			t.Errorf("%d - Expected a ValueError: %T", i+1, err)
		}

	}
}

func TestHelperByteArray(t *testing.T) {

	tests := []testByteData{
		testByteData{input1: []byte{1, 32, 44, math.MaxInt8, 1}, input2: []byte{1, 32, 44, math.MaxInt8, 1}},
		testByteData{input1: []byte{0}, input2: []byte{0}},
		testByteData{input1: []byte{}, input2: []byte{}},
		testByteData{input1: []byte{120}, input2: []byte{120}},
		testByteData{input1: []byte{3, 43, 0}, input2: []byte{3, 43, 0}},
		testByteData{input1: []byte{2, 3}, input2: []byte{2, 3}},
	}

	for i, test := range tests {
		if err := AssertArraysEqual(ByteArrays{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Errorf("%d - %v", i+1, err.Error())
		}
	}
}

func TestHelperDifferentElementsByteArray(t *testing.T) {
	tests := []testByteData{
		testByteData{input1: []byte{1, 3}, input2: []byte{1, 33}},
		testByteData{input1: []byte{1, 5, 7, 8}, input2: []byte{1, 5, 3, 8}},
	}

	for i, test := range tests {
		err := AssertArraysEqual(ByteArrays{Expected: test.input1, Actual: test.input2})

		if err == nil {
			t.Errorf("%d - Expected Exception! Array with different values", i+1)
		}

		_, isValueError := err.(*ValueError)

		if !isValueError {
			t.Errorf("%d - Expected a ValueError: %T", i+1, err)
		}
	}
}

func TestHelperDataArray(t *testing.T) {

	tests := []testData{
		testData{input1: []interface{}{1, 32, 44, math.MaxInt32, 1}, input2: []interface{}{1, 32, 44, math.MaxInt32, 1}},
		testData{input1: []interface{}{true, false, "cat", 3}, input2: []interface{}{true, false, "cat", 3}},
		testData{input1: []interface{}{"AB", "c", "dog"}, input2: []interface{}{"AB", "c", "dog"}},
	}

	for i, test := range tests {
		if err := AssertArraysEqual(DataArrays{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Errorf("%d - %v", i+1, err.Error())
		}
	}
}

func TestHelperDifferentElementsDataArray(t *testing.T) {

	tests := []testData{
		testData{input1: []interface{}{1, 3}, input2: []interface{}{1, 33}},
		testData{input1: []interface{}{1, "5", 7, false}, input2: []interface{}{1, "5", 3, true}},
	}

	for i, test := range tests {
		err := AssertArraysEqual(DataArrays{Expected: test.input1, Actual: test.input2})

		if err == nil {
			t.Errorf("%d - Expected Exception! Array with different values", i+1)
		}

		_, isValueError := err.(*ValueError)

		if !isValueError {
			t.Errorf("%d - Expected a ValueError: %T", i+1, err)
		}
	}
}

func TestHelperDifferentLengthArray(t *testing.T) {

	tests := []testData{
		testData{input1: []interface{}{-1, 3}, input2: []interface{}{1, 3, 4}},
		testData{input1: []interface{}{1, "5", 7, false}, input2: []interface{}{1, "5"}},
	}

	for i, test := range tests {

		err := AssertArraysEqual(DataArrays{Expected: test.input1, Actual: test.input2})
		if err == nil {
			t.Errorf("%d - Expected Exception! Array with different values", i+1)
		}
		_, isLengthError := err.(*LengthError)

		if !isLengthError {
			t.Errorf("%d - Expected a LengthError: %T", i+1, err)
		}
	}
}

func TestHelperGetFuncName(t *testing.T) {
	expectedValue := "AssertArraysEqual"
	nameFunc := GetFuncName(AssertArraysEqual)
	if nameFunc != expectedValue {
		t.Errorf("Expected '%s' - Actual '%s'", expectedValue, nameFunc)
	}
}
