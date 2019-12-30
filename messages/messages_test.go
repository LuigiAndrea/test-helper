package messages

import (
	"fmt"
	"testing"
)

type testMessageError struct {
	input1, input2 interface{}
	output         string
}

func TestHelperGetFuncName(t *testing.T) {
	expectedValue := "ErrorMessage"
	nameFunc := GetFuncName(ErrorMessage)
	if nameFunc != expectedValue {
		t.Error(ErrorMessage(expectedValue, nameFunc))
	}
}

func TestMessageError(t *testing.T) {
	int1, int2 := 5, 12
	string1, string2 := "casa", "home"
	outputMessage := "Expected '%v' - Actual '%v'"

	tests := []testMessageError{
		testMessageError{input1: int1, input2: int2, output: fmt.Sprintf(outputMessage, int1, int2)},
		testMessageError{input1: string1, input2: string2, output: fmt.Sprintf(outputMessage, string1, string2)},
	}

	for _, test := range tests {
		if msg := ErrorMessage(test.input1, test.input2); msg != test.output {
			t.Errorf(fmt.Sprintf(outputMessage, test.output, msg))
		}
	}

}
