package readfile

import (
	"io/ioutil"
)

// ReadFile - struct for store io reader
type ReadFile struct {
	Reader func() ([]byte, error)
}

// RealReader - real reader for read file
func RealReader() ([]byte, error) {
	rawBytes, err := ioutil.ReadFile("../../test-io")
	return rawBytes, err
}

// ReadFileString - main function for read file and return string
func (readFile ReadFile) ReadFileString() string {
	rawBytes, err := readFile.Reader()
	if err != nil {
		return err.Error()
	}
	return string(rawBytes)
}
