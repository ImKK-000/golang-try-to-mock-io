package main

import (
	"fmt"
	"io/ioutil"
)

type ReadFile struct {
	Reader func() (string, error)
}

func FakeIOReadFile() (string, error) {
	return "hi kk!", nil
}

func RealIOReadFile() (string, error) {
	rawBytes, err := ioutil.ReadFile("test-io")
	return string(rawBytes), err
}

func main() {
	fakeReadFile := ReadFile{
		Reader: FakeIOReadFile,
	}
	fmt.Println(fakeReadFile.Reader())

	realReadFile := ReadFile{
		Reader: RealIOReadFile,
	}
	fmt.Println(realReadFile.Reader())
}
