package readfile_test

import (
	"testing"
	"try-to-mock-io/model"
	. "try-to-mock-io/readfile"
)

func Test_ReadFile_Input_IO_Should_Be_Hi_Space_KK_Exclamation_Mark(t *testing.T) {
	expectedResult := "hi kk!"
	mockReadFile := model.ReadFile{
		Reader: MockReaderSuccess,
	}
	actualResult := ReadFile(mockReadFile)

	if expectedResult != actualResult {
		t.Errorf("expect '%s' but it got '%s'", expectedResult, actualResult)
	}
}

func Test_ReadFile_Input_IO_Should_Be_Empty_Message(t *testing.T) {
	expectedResult := ""
	mockReadFile := model.ReadFile{
		Reader: MockReaderError,
	}
	actualResult := ReadFile(mockReadFile)

	if expectedResult != actualResult {
		t.Errorf("expect '%s' but it got '%s'", expectedResult, actualResult)
	}
}
