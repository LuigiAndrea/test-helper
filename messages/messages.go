// Package messages exports standard messages to use in formatting testing outputs
package messages

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

//GetFuncName returns the name of the function passed as parameter
func GetFuncName(f interface{}) string {
	link := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	return link[len(link)-1]
}

//ErrorMessage returns a standard message when actual and expected values are different
func ErrorMessage(ExpectedValue, ActualValue interface{}) string {
	return fmt.Sprintf("Expected: '%#v' - Actual: '%#v'", ExpectedValue, ActualValue)
}

//ErrorMessageTestCount returns a standard formatted message with test count when actual and expected values are different
func ErrorMessageTestCount(n int, testError interface{}) string {
	return fmt.Sprintf("\nTest %d: %v", n, testError)
}
