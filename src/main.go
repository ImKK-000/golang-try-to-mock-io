package main

import (
	"fmt"
	"try-to-mock-io/readfile"
)

func main() {
	r := readfile.ReadFile{
		Reader: readfile.RealReader,
	}
	fmt.Println(r.ReadFileString("../test-io"))
}
