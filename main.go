package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
