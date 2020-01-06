package messages

import "fmt"

func ExampleGetFuncName() {
	expectedValue := "ErrorMessage"
	nameFunc := GetFuncName(ErrorMessage)
	if nameFunc != expectedValue {
		fmt.Println(ErrorMessage(expectedValue, nameFunc))
	}
}
