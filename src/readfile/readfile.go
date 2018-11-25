package readfile

import (
	"io/ioutil"
)

// ReadFile - struct for store io reader
type ReadFile struct {
	Reader func(string) ([]byte, error)
}

// RealReader - real reader for read file
func RealReader(path string) ([]byte, error) {
	rawBytes, err := ioutil.ReadFile(path)
	return rawBytes, err
}

// ReadFileString - main function for read file and return string
func (readFile ReadFile) ReadFileString(path string) string {
	rawBytes, err := readFile.Reader(path)
	if err != nil {
		return err.Error()
	}
	return string(rawBytes)
}
