package readfile_test

import (
	"testing"
	"try-to-mock-io/model"
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
	mockReadFile := model.ReadFile{
		Reader: MockReaderSuccess,
	}
	actualResult := ReadFile(mockReadFile)

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ReadFile_Input_IO_Should_Be_Empty_Message(t *testing.T) {
	expectedResult := ""
	mockReadFile := model.ReadFile{
		Reader: MockReaderError,
	}
	actualResult := ReadFile(mockReadFile)

	assert.Equal(t, expectedResult, actualResult)
}
