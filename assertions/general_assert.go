// Package assertions is used for comparing golang objects
package assertions

import (
	"errors"
	"fmt"
	"reflect"

	m "github.com/LuigiAndrea/test-helper/messages"
)

//AssertDeepEqual checks if two objects are deep equal
func AssertDeepEqual(expected, actual interface{}) error {
	if res := reflect.DeepEqual(expected, actual); !res {
		return errors.New(m.ErrorMessage(expected, actual))
	}

	return nil
}

//AssertNotDeepEqual checks if two objects are not deep equal
func AssertNotDeepEqual(notExpected, actual interface{}) error {
	if res := reflect.DeepEqual(notExpected, actual); res {
		return errors.New("Same objects")
	}

	return nil
}

//ExceptionError records an error when two exceptions are of different type
type ExceptionError struct {
	ExpectedException, CurrentException error
}

func (e *ExceptionError) Error() string {
	return m.ErrorMessage(e.ExpectedException, e.CurrentException)
}

//AssertException checks if the exception fired has right type
func AssertException(expectedException, currentException error) error {
	if res := reflect.TypeOf(expectedException) != reflect.TypeOf(currentException); res {
		return &ExceptionError{ExpectedException: expectedException, CurrentException: currentException}
	}

	return nil
}

//AssertDeepException checks if the exceptions have same values and types
func AssertDeepException(expectedException, currentException error) error {

	if !errors.Is(currentException, expectedException) {
		return &ExceptionError{ExpectedException: expectedException, CurrentException: currentException}
	}

	return nil
}

//AssertPanic checks if the function passed as parameter panic
func AssertPanic(f func()) (e error) {
	defer func() error {
		if r := recover(); r == nil {
			e = fmt.Errorf("Expected '%s' to panic", m.GetFuncName(f))
		}
		return nil
	}()

	f()

	return nil
}
