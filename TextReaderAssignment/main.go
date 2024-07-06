package main

import (
	"io"
	"os"
)

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)
	if err != nil {
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
