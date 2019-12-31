package assertions

import (
	"errors"
	"fmt"

	m "github.com/LuigiAndrea/test-helper/messages"
)

//AssertSlicesEqual Check if two slices have the same values and in the same order
func AssertSlicesEqual(slices CheckSlices) error {
	if !slices.SameLength() {
		return &LengthError{Err: errors.New("Slices with different length")}
	}

	for i := 0; i < slices.Size(); i++ {
		if v := slices.AreEqual(i); !v {
			return slices.GetError(i)
		}
	}

	return nil
}

//CheckSlices interface is used for comparing two slices in testing
type CheckSlices interface {
	SameLength() bool
	AreEqual(i int) bool
	Size() int
	GetError(i int) error
}

//LengthError records an error when two slices have different lengths
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

// StringSlicesMatch attaches the methods of CheckSlices to struct StringSlicesMatch
type StringSlicesMatch struct {
	Expected []string
	Actual   []string
}

// SameLength checks if the two slices have the same length
func (s StringSlicesMatch) SameLength() bool { return len(s.Expected) == len(s.Actual) }

// AreEqual checks if the two slices have the same value at position i
func (s StringSlicesMatch) AreEqual(i int) bool { return s.Expected[i] == s.Actual[i] }

// Size is the length of StringSlicesMatch
func (s StringSlicesMatch) Size() int { return len(s.Expected) }

// GetError displays an error message when the values at position i are different
func (s StringSlicesMatch) GetError(i int) error {
	return &ValueError{x: s.Expected[i], y: s.Actual[i], pos: i}
}

// IntSlicesMatch attaches the methods of CheckSlices to struct IntSlicesMatch
type IntSlicesMatch struct {
	Expected []int
	Actual   []int
}

// SameLength checks if the two slices have the same length
func (islice IntSlicesMatch) SameLength() bool { return len(islice.Expected) == len(islice.Actual) }

// AreEqual checks if the two slices have the same value at position i
func (islice IntSlicesMatch) AreEqual(i int) bool { return islice.Expected[i] == islice.Actual[i] }

// Size is the length of IntSlicesMatch struct
func (islice IntSlicesMatch) Size() int { return len(islice.Expected) }

// GetError displays an error message when the values at position i are different
func (islice IntSlicesMatch) GetError(i int) error {
	return &ValueError{x: islice.Expected[i], y: islice.Actual[i], pos: i}
}

//Float64SlicesMatch attaches the methods of CheckSlices to struct Float64SlicesMatch
type Float64SlicesMatch struct {
	Expected []float64
	Actual   []float64
}

// SameLength checks if the two slices have the same length
func (f Float64SlicesMatch) SameLength() bool { return len(f.Expected) == len(f.Actual) }

// AreEqual checks if the two slices have the same value at position i
func (f Float64SlicesMatch) AreEqual(i int) bool { return f.Expected[i] == f.Actual[i] }

// Size is the length of Float64SlicesMatch struct
func (f Float64SlicesMatch) Size() int { return len(f.Expected) }

// GetError displays an error message when the values at position i are different
func (f Float64SlicesMatch) GetError(i int) error {
	return &ValueError{x: f.Expected[i], y: f.Actual[i], pos: i}
}

//ByteSlicesMatch attaches the methods of CheckSlices to struct ByteSlicesMatch
type ByteSlicesMatch struct {
	Expected []byte
	Actual   []byte
}

// SameLength checks if the two slices have the same length
func (b ByteSlicesMatch) SameLength() bool { return len(b.Expected) == len(b.Actual) }

// AreEqual checks if the two slices have the same value at position i
func (b ByteSlicesMatch) AreEqual(i int) bool { return b.Expected[i] == b.Actual[i] }

// Size is the length of ByteSlicesMatch struct
func (b ByteSlicesMatch) Size() int { return len(b.Expected) }

// GetError displays an error message when the values at position i are different
func (b ByteSlicesMatch) GetError(i int) error {
	return &ValueError{x: b.Expected[i], y: b.Actual[i], pos: i}
}

//DataSlicesMatch attaches the methods of CheckSlices to struct DataSlicesMatch
type DataSlicesMatch struct {
	Expected []interface{}
	Actual   []interface{}
}

// SameLength checks if the two slices have the same length
func (d DataSlicesMatch) SameLength() bool { return len(d.Expected) == len(d.Actual) }

// AreEqual checks if the two slices have the same value at position i
func (d DataSlicesMatch) AreEqual(i int) bool { return d.Expected[i] == d.Actual[i] }

// Size is the length of DataSlicesMatch struct
func (d DataSlicesMatch) Size() int { return len(d.Expected) }

// GetError displays an error message when the values at position i are different
func (d DataSlicesMatch) GetError(i int) error {
	return &ValueError{x: d.Expected[i], y: d.Actual[i], pos: i}
}
