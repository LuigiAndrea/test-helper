package assertions

import (
	"errors"
	"reflect"

	m "github.com/LuigiAndrea/test-helper/messages"
)

//AssertDeepEqual Check if two objects are deep equal
func AssertDeepEqual(expected interface{}, actual interface{}) error {
	if res := reflect.DeepEqual(expected, actual); !res {
		return errors.New(m.ErrorMessage(expected, actual))
	}

	return nil
}

//ExceptionError records an error when two exceptions are of different type
type ExceptionError struct {
	ExpectedException, CurrentException interface{}
}

func (e *ExceptionError) Error() string {
	return m.ErrorMessage(e.ExpectedException, e.CurrentException)
}

//AssertException Check if the exception fired is right
func AssertException(expectedException interface{}, currentException interface{}) error {
	if res := reflect.TypeOf(expectedException) != reflect.TypeOf(currentException); res {
		return &ExceptionError{ExpectedException: expectedException, CurrentException: currentException}
	}
	return nil
}
