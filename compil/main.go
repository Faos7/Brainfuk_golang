package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	fileName := os.Args[1]
	code, err :=ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}
	compiler := NewCompiler(string(code))
	instructions := compiler.Compile()

	m:= NewMachine(instructions, os.Stdin, os.Stdout)
	m.Execute()
}
