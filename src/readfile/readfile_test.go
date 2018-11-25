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
	readFile := ReadFile{
		Reader: MockReaderSuccess,
	}
	actualResult := readFile.ReadFileString()

	assert.Equal(t, expectedResult, actualResult)
}

func Test_ReadFile_Input_IO_Should_Be_Error_Message(t *testing.T) {
	expectedResult := "Reader Error!"
	readFile := ReadFile{
		Reader: MockReaderError,
	}
	actualResult := readFile.ReadFileString()

	assert.Equal(t, expectedResult, actualResult)
}
