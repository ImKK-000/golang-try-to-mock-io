package main

import (
	"fmt"
	"io/ioutil"
)

type ReadFile struct {
	Reader func() ([]byte, error)
}

func FakeIOReadFile() ([]byte, error) {
	return []byte("hi kk!"), nil
}

func RealIOReadFile() ([]byte, error) {
	rawBytes, err := ioutil.ReadFile("../test-io")
	return rawBytes, err
}

func main() {
	fakeReadFile := ReadFile{
		Reader: FakeIOReadFile,
	}
	fakeRawBytes, _ := fakeReadFile.Reader()
	fmt.Println(string(fakeRawBytes))

	realReadFile := ReadFile{
		Reader: RealIOReadFile,
	}
	realRawBytes, _ := realReadFile.Reader()
	fmt.Println(string(realRawBytes))
}
