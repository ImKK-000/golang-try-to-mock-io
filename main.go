package main

import (
	"fmt"
	"io/ioutil"
)

func MockReadFile() string {
	return "hi kk!"
}

func ReadFile() string {
	rawBytes, _ := ioutil.ReadFile("test-io")
	return string(rawBytes)
}

func main() {
	fmt.Println(MockReadFile())
	fmt.Println(ReadFile())
}
