#Brainfuk compiller writen in Go

This is study project.

## How to run:
1. Change directory to src/
2. Compile Machine.go

## For example: 
1. `cd /Brainfuk_golang/compil/`
2. `go build -o machine && time ./machine ./mandelbrot.b >/dev/null`

PS. mandelbrod is benchmark. Running compiller with 'time' options alllows to check how many time PC spends to compile and run benchmark.

## BF files

There are 2 .bf files - helloWorld and mandelbrot

## Specs

There are 2 working projects - interpritor and compiller.

Compiller optimizes code, therefor time to run benchmark using compiller is 2 times less then using interpritor.

This proves that compiller realy optimizes code.
