package main // main package

import "github.com/01-edu/z01" // import z01 package

func main() { // main function
	for i := 'a'; i <= 'z'; i++ { // loop through lowercase letters
		z01.PrintRune(i) // print each letter
	}
	z01.PrintRune('\n') // end line
}
