package readfile

import (
	"io/ioutil"
)

type ReadFile struct {
	Reader func() ([]byte, error)
}

func RealReader() ([]byte, error) {
	rawBytes, err := ioutil.ReadFile("../../test-io")
	return rawBytes, err
}

func ReadFileString(readFile ReadFile) string {
	rawBytes, err := readFile.Reader()
	if err != nil {
		return ""
	}
	return string(rawBytes)
}
