// Package assertions is used for comparing golang objects
package assertions

import (
	"errors"
	"reflect"

	m "github.com/LuigiAndrea/test-helper/messages"
)

//AssertDeepEqual checks if two objects are deep equal
func AssertDeepEqual(expected interface{}, actual interface{}) error {
	if res := reflect.DeepEqual(expected, actual); !res {
		return errors.New(m.ErrorMessage(expected, actual))
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
