package main

import (
	"fmt"
	"os"

	"piscineModule/piscine/globalVariable"
)

func extractingArgs(i int) []int {
	globalVariable.IntoSlice = make([]int, 9)
	for j := 0; j < 9; j++ {
		if os.Args[i][j] >= '0' && os.Args[i][j] <= '9' {
			globalVariable.IntoSlice[j] = int(os.Args[i][j] - '0')
		} else {
			globalVariable.IntoSlice[j] = 0
		}
	}
	return globalVariable.IntoSlice
}

func checkingRules(board [][]int, i, j, nbr int) bool {
	for p := 0; p < 9; p++ {
		if board[i][p] == nbr || board[p][j] == nbr {
			return false
		}
	}

	rowStart := (i / 3) * 3
	colStart := (j / 3) * 3
	for m := rowStart; m < rowStart+3; m++ {
		for n := colStart; n < colStart+3; n++ {
			if board[m][n] == nbr {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for nbr := 1; nbr <= 9; nbr++ {
					if checkingRules(board, i, j, nbr) {
						board[i][j] = nbr
						if solveSudoku(board) {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	board := make([][]int, 9)
	for i := 0; i < 9; i++ {
		board[i] = extractingArgs(i + 1)
	}

	if solveSudoku(board) {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Print(board[i][j], " ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}
}
