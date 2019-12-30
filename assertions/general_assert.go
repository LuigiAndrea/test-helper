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
