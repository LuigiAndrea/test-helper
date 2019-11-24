package gotesthelper

import (
	"math"
	"testing"
)

func TestHelperString(t *testing.T) {
	actualValues := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}
	expectedValues := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}

	for i, actualValue := range actualValues {
		if err := CheckArraySameValues(StringArrays{Expected: expectedValues[i], Actual: actualValue}); err != nil {
			t.Error(err.Error())
			break
		}
	}
}

func TestHelperInt(t *testing.T) {
	actualValues := [][]int{{1, 32, 44322, math.MaxInt64, math.MinInt64}, {133}, {0}, {}, {-3, 43, -0}}
	expectedValues := [][]int{{1, 32, 44322, math.MaxInt64, math.MinInt64}, {133}, {0}, {}, {-3, 43, 0}}

	for i, actualValue := range actualValues {
		if err := CheckArraySameValues(IntArrays{Expected: expectedValues[i], Actual: actualValue}); err != nil {
			t.Error(err.Error())
			break
		}
	}
}

func TestHelperFloat64(t *testing.T) {
	actualValues := [][]float64{{1, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(2)}, {133.0}, {0.0}, {}, {-3.0, 43.0, -0.0}, {2.5, 3.3}}
	expectedValues := [][]float64{{1.0, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(200)}, {133.0}, {0.0}, {}, {-3.0, 43.0, 0.0}, {2.5, 3.3}}

	for i, actualValue := range actualValues {
		if err := CheckArraySameValues(Float64Arrays{Expected: expectedValues[i], Actual: actualValue}); err != nil {
			t.Error(err.Error())
			break
		}
	}
}

func TestHelperByte(t *testing.T) {
	actualValues := [][]byte{{1, 32, 44, math.MaxInt8, 1}, {120}, {0}, {}, {3, 43, -0}, {2, 3}}
	expectedValues := [][]byte{{1, 32, 44, math.MaxInt8, 1}, {120}, {0}, {}, {3, 43, 0}, {2, 3}}

	for i, actualValue := range actualValues {
		if err := CheckArraySameValues(ByteArrays{Expected: expectedValues[i], Actual: actualValue}); err != nil {
			t.Error(err.Error())
			break
		}
	}
}

func TestHelperDataType(t *testing.T) {
	actualValues := [][]interface{}{{-1, 3}, {}, {true, false, false}, {"L", "UI", "GI"}}
	expectedValues := [][]interface{}{{-1, 3}, {}, {true, false, false}, {"L", "UI", "GI"}}

	for i, actualValue := range actualValues {
		if err := CheckArraySameValues(DataArrays{Expected: expectedValues[i], Actual: actualValue}); err != nil {
			t.Error(err.Error())
			break
		}
	}
}

func TestHelperDifferentLengthArray(t *testing.T) {
	arraysOne, arraysTwo := [][]interface{}{{-1, 3}, {"B", 5, 7, 8}}, [][]interface{}{{-1, 3, 4, 6}, {"B"}}

	for i, aOne := range arraysOne {

		err := CheckArraySameValues(DataArrays{Expected: arraysTwo[i], Actual: aOne})
		if err == nil {
			t.Error("Expected Exception! Array with different lengths")
			break
		}
		_, isLengthError := err.(*LengthError)

		if !isLengthError {
			t.Error("Expected a LengthError")
			break
		}
	}
}

func TestHelperDifferentElementsDataArray(t *testing.T) {
	actualValues, expectedValues := [][]interface{}{{-1, 33}, {1, 5, 7, "8"}}, [][]interface{}{{-1, "33"}, {1, 5, 3, 8}}

	for i, expectedValue := range expectedValues {
		if err := CheckArraySameValues(DataArrays{Expected: expectedValue, Actual: actualValues[i]}); err == nil {
			t.Error("Expected Exception! Array with different values")
			break
		}
	}
}

func TestHelperDifferentElementsIntArray(t *testing.T) {
	actualValues, expectedValues := [][]int{{-1, 3}, {1, 5, 7, 8}}, [][]int{{-1, 33}, {1, 5, 3, 8}}

	for i, expectedValue := range expectedValues {
		if err := CheckArraySameValues(IntArrays{Expected: expectedValue, Actual: actualValues[i]}); err == nil {
			t.Error("Expected Exception! Array with different values")
			break
		}
	}
}

func TestHelperDifferentElementsFloatArray(t *testing.T) {
	actualValues, expectedValues := [][]float64{{-1, 3}, {1, 5, 7, 8}}, [][]float64{{-1, 33}, {1, 5, 3, 8}}

	for i, expectedValue := range expectedValues {
		if err := CheckArraySameValues(Float64Arrays{Expected: expectedValue, Actual: actualValues[i]}); err == nil {
			t.Error("Expected Exception! Array with different values")
			break
		}
	}
}

func TestHelperDifferentElementsStringArray(t *testing.T) {
	actualValues, expectedValues := [][]string{{"-1", "1"}, {"1", "5", "7", "8"}}, [][]string{{"-1", "13"}, {"1", "5", "3", "8"}}

	for i, expectedValue := range expectedValues {
		if err := CheckArraySameValues(StringArrays{Expected: expectedValue, Actual: actualValues[i]}); err == nil {
			t.Error("Expected Exception! Array with different values")
			break
		}
	}
}

func TestHelperDifferentElementsByteArray(t *testing.T) {
	actualValues, expectedValues := [][]byte{{1, 3}, {1, 5, 7, 8}}, [][]byte{{1, 33}, {1, 5, 3, 8}}

	for i, expectedValue := range expectedValues {
		if err := CheckArraySameValues(ByteArrays{Expected: expectedValue, Actual: actualValues[i]}); err == nil {
			t.Error("Expected Exception! Array with different values")
			break
		}
	}
}

func TestHelperGetFuncName(t *testing.T) {
	expectedValue := "CheckArraySameValues"
	nameFunc := GetFuncName(CheckArraySameValues)
	if nameFunc != expectedValue {
		t.Errorf("\nExpected '%s' - Actual '%s'", expectedValue, nameFunc)
	}
}
