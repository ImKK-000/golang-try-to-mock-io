package main

import (
	"flag"
	"fmt"
	"try-to-mock-io/readfile"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "../test-io", "set filename")
	flag.Parse()

	r := readfile.ReadFile{
		Reader: readfile.RealReader,
	}
	fmt.Println(r.ReadFileString(filename))
}
