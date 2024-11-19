package main

import ("github.com/01-edu/z01"
      "os"
)
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if (nbr) == 1  {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg :=len(os.Args[1:])
	if isEven(lengthOfArg) {
		EvenMsg := "I have an even number of arguments"
		printStr(EvenMsg)
	} else {
		OddMsg := "I have an odd number of arguments"
		printStr(OddMsg)
	}
}