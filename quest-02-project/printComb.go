package piscine

import "github.com/01-edu/z01"

func printComb() { // functuon
	var i, j, k int          // initialize variables
	for i = 0; i <= 7; i++ { // loop 1st nuber set
		for j = i + 1; j <= 8; j++ { // loop 2nd number set
			for k = j + 1; k <= 9; k++ { // loop 3rd number set
				z01.PrintRune('0' + rune(i)) // print 1st number
				z01.PrintRune('0' + rune(j)) // print 2nd number
				z01.PrintRune('0' + rune(k)) // print 3rd number

				if i != 7 || j != 8 || k != 9 { // stop
					z01.PrintRune(',') // print comma
					z01.PrintRune(' ') // print space after each number and comma

				}
			}
		}
	}
	z01.PrintRune('\n') // end line
}

