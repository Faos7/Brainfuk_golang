package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
	m := NewMachine(string(code), os.Stdin, os.Stdout)
	m.Execute()
}
