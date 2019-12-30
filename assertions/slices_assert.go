package assertions

import (
	"errors"
	"fmt"

	m "github.com/LuigiAndrea/test-helper/messages"
)

//AssertArraysEqual Check if two arrays have the same values and in the same order
func AssertArraysEqual(arrays CheckArrays) error {
	if !arrays.SameLength() {
		return &LengthError{Err: errors.New("Arrays with different length")}
	}

	for i := 0; i < arrays.Size(); i++ {
		if v := arrays.AreEqual(i); !v {
			return arrays.GetError(i)
		}
	}

	return nil
}

//CheckArrays interface is used for comparing two arrays in testing
type CheckArrays interface {
	SameLength() bool
	AreEqual(i int) bool
	Size() int
	GetError(i int) error
}

//LengthError records an error when two slice/array have different length
type LengthError struct {
	Err error
}

func (e *LengthError) Error() string {
	return e.Err.Error()
}

//ValueError records an error when two elements have different values
type ValueError struct {
	x, y interface{}
	pos  int
}

func (e *ValueError) Error() string {
	return fmt.Sprintf("%s at position %d", m.ErrorMessage(e.x, e.y), e.pos)
}

// StringArrays attaches the methods of CheckArrays to struct StringArrays
type StringArrays struct {
	Expected []string
	Actual   []string
}

// SameLength checks if the two arrays have the same length
func (s StringArrays) SameLength() bool { return len(s.Expected) == len(s.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (s StringArrays) AreEqual(i int) bool { return s.Expected[i] == s.Actual[i] }

// Size is the length of StringArrays
func (s StringArrays) Size() int { return len(s.Expected) }

// GetError displays an error message when the values at position i are different
func (s StringArrays) GetError(i int) error {
	return &ValueError{x: s.Expected[i], y: s.Actual[i], pos: i}
}

// IntArrays attaches the methods of CheckArrays to struct IntArrays
type IntArrays struct {
	Expected []int
	Actual   []int
}

// SameLength checks if the two arrays have the same length
func (iarray IntArrays) SameLength() bool { return len(iarray.Expected) == len(iarray.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (iarray IntArrays) AreEqual(i int) bool { return iarray.Expected[i] == iarray.Actual[i] }

// Size is the length of IntArrays struct
func (iarray IntArrays) Size() int { return len(iarray.Expected) }

// GetError displays an error message when the values at position i are different
func (iarray IntArrays) GetError(i int) error {
	return &ValueError{x: iarray.Expected[i], y: iarray.Actual[i], pos: i}
}

//Float64Arrays attaches the methods of CheckArrays to struct Float64Arrays
type Float64Arrays struct {
	Expected []float64
	Actual   []float64
}

// SameLength checks if the two arrays have the same length
func (f Float64Arrays) SameLength() bool { return len(f.Expected) == len(f.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (f Float64Arrays) AreEqual(i int) bool { return f.Expected[i] == f.Actual[i] }

// Size is the length of Float64Arrays struct
func (f Float64Arrays) Size() int { return len(f.Expected) }

// GetError displays an error message when the values at position i are different
func (f Float64Arrays) GetError(i int) error {
	return &ValueError{x: f.Expected[i], y: f.Actual[i], pos: i}
}

//ByteArrays attaches the methods of CheckArrays to struct ByteArrays
type ByteArrays struct {
	Expected []byte
	Actual   []byte
}

// SameLength checks if the two arrays have the same length
func (b ByteArrays) SameLength() bool { return len(b.Expected) == len(b.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (b ByteArrays) AreEqual(i int) bool { return b.Expected[i] == b.Actual[i] }

// Size is the length of ByteArrays struct
func (b ByteArrays) Size() int { return len(b.Expected) }

// GetError displays an error message when the values at position i are different
func (b ByteArrays) GetError(i int) error {
	return &ValueError{x: b.Expected[i], y: b.Actual[i], pos: i}
}

//DataArrays attaches the methods of CheckArrays to struct DataArrays
type DataArrays struct {
	Expected []interface{}
	Actual   []interface{}
}

// SameLength checks if the two arrays have the same length
func (d DataArrays) SameLength() bool { return len(d.Expected) == len(d.Actual) }

// AreEqual checks if the two arrays have the same value at position i
func (d DataArrays) AreEqual(i int) bool { return d.Expected[i] == d.Actual[i] }

// Size is the length of DataArrays struct
func (d DataArrays) Size() int { return len(d.Expected) }

// GetError displays an error message when the values at position i are different
func (d DataArrays) GetError(i int) error {
	return &ValueError{x: d.Expected[i], y: d.Actual[i], pos: i}
}