package main

import "github.com/01-edu/z01"

func main() {
	for i := 'z'; i >= 'a'; i-- { // loop through letters in reverse order
		z01.PrintRune(i) // print each letter
	}
	z01.PrintRune('\n')
}
