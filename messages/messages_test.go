package messages

import (
	"fmt"
	"testing"
)

type testErrorMessage struct {
	input1, input2 interface{}
	output         string
}

type testErrorMessageCount struct {
	input1 int
	input2 interface{}
	output string
}

func TestHelperGetFuncName(t *testing.T) {
	expectedValue := "ErrorMessage"
	nameFunc := GetFuncName(ErrorMessage)
	if nameFunc != expectedValue {
		t.Error(ErrorMessage(expectedValue, nameFunc))
	}
}

func TestErrorMessage(t *testing.T) {
	int1, int2 := 5, 12
	string1, string2 := "casa", "home"
	outputMessage := "Expected: '%#v' - Actual: '%#v'"

	tests := []testErrorMessage{
		{input1: int1, input2: int2, output: fmt.Sprintf(outputMessage, int1, int2)},
		{input1: string1, input2: string2, output: fmt.Sprintf(outputMessage, string1, string2)},
	}

	for i, test := range tests {
		if msg := ErrorMessage(test.input1, test.input2); msg != test.output {
			t.Error(ErrorMessageTestCount(i+1, fmt.Sprintf(outputMessage, test.output, msg)))
		}
	}
}

func TestErrorMessageCount(t *testing.T) {
	testCount1, testCount2 := 1, 2
	testToPrint1, testToPrint2 := "Message to visualize", "Message to visualize 2"
	outputMessage := "\nTest %d: %s"

	tests := []testErrorMessageCount{
		{input1: testCount1, input2: testToPrint1, output: fmt.Sprintf(outputMessage, testCount1, testToPrint1)},
		{input1: testCount2, input2: testToPrint2, output: fmt.Sprintf(outputMessage, testCount2, testToPrint2)},
	}

	for _, test := range tests {

		if msg := ErrorMessageTestCount(test.input1, test.input2); msg != test.output {
			t.Errorf(ErrorMessage(test.output, msg))
		}
	}
}
