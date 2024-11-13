package main

import "github.com/01-edu/z01"

func main() {
	for h := '0'; h <= '9'; h++ {
		z01.PrintRune(h)
	}
	z01.PrintRune('\n')
}
