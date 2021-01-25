package assertions

import (
	"math"
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
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

type testBoolData struct {
	input1, input2 []bool
}

type testData struct {
	input1, input2 []interface{}
}

func TestHelperStringSlices(t *testing.T) {
	tests := []testStringData{
		{input1: []string{"a", "b", "c", "d", "e"}, input2: []string{"a", "b", "c", "d", "e"}},
		{input1: []string{"c"}, input2: []string{"c"}},
		{input1: []string{""}, input2: []string{""}},
		{input1: []string{}, input2: []string{}},
		{input1: []string{"d", "e", "a"}, input2: []string{"d", "e", "a"}},
		{input1: []string{"T", "R", "E"}, input2: []string{"T", "R", "E"}},
		{input1: []string{"CC", "AAA", "BB"}, input2: []string{"CC", "AAA", "BB"}},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(StringSlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestHelperDifferentElementsStringSlices(t *testing.T) {
	var ve *ValueError
	tests := []testStringData{
		{input1: []string{"-1", "1"}, input2: []string{"-1", "13"}},
		{input1: []string{"1", "5", "7", "8"}, input2: []string{"1", "5", "3", "8"}},
		{input1: []string{""}, input2: []string{" "}},
	}

	for i, test := range tests {
		err := AssertSlicesEqual(StringSlicesMatch{Expected: test.input1, Actual: test.input2})
		if errExc := AssertException(ve, err); errExc != nil {
			t.Error(m.ErrorMessageTestCount(i+1, errExc.Error()))
		}
	}
}

func TestHelperIntSlices(t *testing.T) {
	tests := []testIntData{
		{input1: []int{1, 32, 44322, math.MaxInt64, math.MinInt64}, input2: []int{1, 32, 44322, math.MaxInt64, math.MinInt64}},
		{input1: []int{133}, input2: []int{133}},
		{input1: []int{0}, input2: []int{0}},
		{input1: []int{}, input2: []int{}},
		{input1: []int{-3, 43, -0}, input2: []int{-3, 43, -0}},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(IntSlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestHelperDifferentElementsIntSlices(t *testing.T) {
	var ve *ValueError
	tests := []testIntData{
		{input1: []int{-1, 3}, input2: []int{-1, 33}},
		{input1: []int{1, 5, 7, 8}, input2: []int{1, 5, 3, 8}},
	}

	for i, test := range tests {
		err := AssertSlicesEqual(IntSlicesMatch{Expected: test.input1, Actual: test.input2})
		if errExc := AssertException(ve, err); errExc != nil {
			t.Error(m.ErrorMessageTestCount(i+1, errExc.Error()))
		}
	}
}

func TestHelperFloat64Slices(t *testing.T) {

	tests := []testFloatData{
		{input1: []float64{1, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(2)}, input2: []float64{1, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(2)}},
		{input1: []float64{133}, input2: []float64{133}},
		{input1: []float64{0.0}, input2: []float64{0.0}},
		{input1: []float64{}, input2: []float64{}},
		{input1: []float64{-3.0, 43.0, -0.0}, input2: []float64{-3.0, 43.0, -0.0}},
		{input1: []float64{2.5, 3.3}, input2: []float64{2.5, 3.3}},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(Float64SlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestHelperDifferentElementsFloatSlices(t *testing.T) {
	var ve *ValueError
	tests := []testFloatData{
		{input1: []float64{-1, 3}, input2: []float64{-1, 33}},
		{input1: []float64{1, 5, 7, 8}, input2: []float64{1, 5, 3, 8}},
	}

	for i, test := range tests {
		err := AssertSlicesEqual(Float64SlicesMatch{Expected: test.input1, Actual: test.input2})
		if errExc := AssertException(ve, err); errExc != nil {
			t.Error(m.ErrorMessageTestCount(i+1, errExc.Error()))
		}
	}
}

func TestHelperByteSlices(t *testing.T) {

	tests := []testByteData{
		{input1: []byte{1, 32, 44, math.MaxInt8, 1}, input2: []byte{1, 32, 44, math.MaxInt8, 1}},
		{input1: []byte{0}, input2: []byte{0}},
		{input1: []byte{}, input2: []byte{}},
		{input1: []byte{120}, input2: []byte{120}},
		{input1: []byte{3, 43, 0}, input2: []byte{3, 43, 0}},
		{input1: []byte{2, 3}, input2: []byte{2, 3}},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(ByteSlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestHelperDifferentElementsByteSlices(t *testing.T) {
	var ve *ValueError
	tests := []testByteData{
		{input1: []byte{1, 3}, input2: []byte{1, 33}},
		{input1: []byte{1, 5, 7, 8}, input2: []byte{1, 5, 3, 8}},
	}

	for i, test := range tests {
		err := AssertSlicesEqual(ByteSlicesMatch{Expected: test.input1, Actual: test.input2})

		if errExc := AssertException(ve, err); errExc != nil {
			t.Error(m.ErrorMessageTestCount(i+1, errExc.Error()))
		}
	}
}

func TestHelperBoolSlices(t *testing.T) {

	tests := []testBoolData{
		{input1: []bool{true, false, true}, input2: []bool{true, false, true}},
		{input1: []bool{false}, input2: []bool{false}},
		{input1: []bool{}, input2: []bool{}},
		{input1: nil, input2: nil},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(BoolSlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestHelperDifferentElementsBoolSlices(t *testing.T) {
	var ve *ValueError
	tests := []testBoolData{
		{input1: []bool{true, true}, input2: []bool{true, false}},
		{input1: []bool{false}, input2: []bool{true}},
	}

	for i, test := range tests {
		err := AssertSlicesEqual(BoolSlicesMatch{Expected: test.input1, Actual: test.input2})

		if errExc := AssertException(ve, err); errExc != nil {
			t.Error(m.ErrorMessageTestCount(i+1, errExc.Error()))
		}
	}
}

func TestHelperDataSlices(t *testing.T) {

	tests := []testData{
		{input1: []interface{}{1, 32, 44, math.MaxInt32, 1}, input2: []interface{}{1, 32, 44, math.MaxInt32, 1}},
		{input1: []interface{}{true, false, "cat", 3}, input2: []interface{}{true, false, "cat", 3}},
		{input1: []interface{}{"AB", "c", "dog"}, input2: []interface{}{"AB", "c", "dog"}},
	}

	for i, test := range tests {
		if err := AssertSlicesEqual(DataSlicesMatch{Expected: test.input1, Actual: test.input2}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err.Error()))
		}
	}
}

func TestHelperDifferentElementsDataSlices(t *testing.T) {
	var ve *ValueError
	tests := []testData{
		{input1: []interface{}{1, 3}, input2: []interface{}{1, 33}},
		{input1: []interface{}{1, "5", 7, false}, input2: []interface{}{1, "5", 3, true}},
		{input1: []interface{}{1, "5", 7, false}, input2: []interface{}{"5", 3, true, 12}},
	}

	for i, test := range tests {
		err := AssertSlicesEqual(DataSlicesMatch{Expected: test.input1, Actual: test.input2})

		if errExc := AssertException(ve, err); errExc != nil {
			t.Error(m.ErrorMessageTestCount(i+1, errExc.Error()))
		} else if len(err.Error()) == 0 {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected to receive an error message"))
		}
	}
}

func TestHelperDifferentLengthSlices(t *testing.T) {
	var le *LengthError
	tests := []testData{
		{input1: []interface{}{-1, 3}, input2: []interface{}{1, 3, 4}},
		{input1: []interface{}{1, "5", 7, false}, input2: []interface{}{1, "5"}},
	}

	for i, test := range tests {

		err := AssertSlicesEqual(DataSlicesMatch{Expected: test.input1, Actual: test.input2})
		if errExc := AssertException(le, err); errExc != nil {
			t.Error(m.ErrorMessageTestCount(i+1, errExc.Error()))
		} else if len(err.Error()) == 0 {
			t.Error(m.ErrorMessageTestCount(i+1, "Expected to receive an error message"))
		}
	}
}
