package readfile_test

import (
	"testing"
	. "try-to-mock-io/readfile"

	"github.com/stretchr/testify/assert"
)

func Test_RealReader_Should_Be_Bytes_And_No_Error(t *testing.T) {
	expectedBytesResult := []byte("hi kk!")
	expectedErrorResult := error(nil)

	actualBytesResult, actualErrorResult := RealReader()

	assert.Equal(t, expectedBytesResult, actualBytesResult)
	assert.Equal(t, expectedErrorResult, actualErrorResult)
}

func Test_ReadFile_Input_IO_Should_Be_Hi_Space_KK_Exclamation_Mark(t *testing.T) {
	expectedResult := "hi kk!"
	mockReadFile := ReadFile{
		Reader: MockReaderSuccess,
	}
	actualResult := ReadFileString(mockReadFile)

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ReadFile_Input_IO_Should_Be_Empty_Message(t *testing.T) {
	expectedResult := ""
	mockReadFile := ReadFile{
		Reader: MockReaderError,
	}
	actualResult := ReadFileString(mockReadFile)

	assert.Equal(t, expectedResult, actualResult)
}
