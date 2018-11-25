package readfile

import (
	"io/ioutil"
	"try-to-mock-io/model"
)

func RealReader() ([]byte, error) {
	rawBytes, err := ioutil.ReadFile("../../test-io")
	return rawBytes, err
}

func ReadFile(readFile model.ReadFile) string {
	rawBytes, err := readFile.Reader()
	if err != nil {
		return ""
	}
	return string(rawBytes)
}
