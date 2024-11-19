package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := range os.Args[1:] {
		arg := os.Args[len(os.Args)-i-1]
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
