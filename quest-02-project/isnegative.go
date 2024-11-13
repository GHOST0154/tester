package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) { // Function to check if a number is negative
	if nb < 0 { // If the number is negative
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
