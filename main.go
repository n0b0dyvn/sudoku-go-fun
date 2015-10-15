package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

const UNASSIGNED = 0

func PrintGrid(S [][]int) {
	for i := range S {
		if i%3 == 0 {
			fmt.Println("- - - - - - - - - - - - - - - - ")
		}
		fmt.Print("|")
		for j := range S[i] {
			fmt.Print(fmt.Sprintf(" %v ", S[i][j]))
			if j%3 == 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i == 8 {
			fmt.Println("- - - - - - - - - - - - - - - - ")
		}
	}
}

func findUnassigned(S [][]int) (int, int, bool) {
	for i := range S {
		for j := range S[i] {
			if S[i][j] == UNASSIGNED {
				return i, j, true
			}
		}
	}
	return 9, 9, false
}

func UsedInRow(S [][]int, row, num int) bool {
	for j := range S[row] {
		if S[row][j] == num {
			return true
		}
	}
	return false
}

func UsedInCol(S [][]int, col, num int) bool {
	for i := range S {
		if S[i][col] == num {
			return true
		}
	}
	return false
}

func UsedInBox(S [][]int, row, col, num int) bool {
	row_start := (row / 3) * 3
	col_start := (col / 3) * 3
	for i := row_start; i < row_start+3; i++ {
		for j := col_start; j < col_start+3; j++ {
			if S[i][j] == num {
				return true
			}
		}
	}
	return false
}

func isSafe(sudoku [][]int, row, col, num int) bool {
	return !UsedInRow(sudoku, row, num) && !UsedInCol(sudoku, col, num) && !UsedInBox(sudoku, row, col, num)
}

func Solve(sudoku [][]int) bool {
	row, col, CheckUnAssign := findUnassigned(sudoku)
	if !CheckUnAssign {
		return true
	}
	for num := 1; num <= 9; num++ {
		if isSafe(sudoku, row, col, num) {
			sudoku[row][col] = num
			if Solve(sudoku) {
				return true
			}
			sudoku[row][col] = UNASSIGNED
		}
	}
	return false
}

func main() {
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
	}
	// Input Sudoku
	f := flag.String("f", "sudoku.txt", "file input")
	b, _ := ioutil.ReadFile(*f)
	row := 0
	col := 0
	for _, v := range b {
		if v >= 48 && v <= 57 {
			sudoku[row][col] = int(v) - 48
			col++
			if col == 9 {
				col = 0
				row++
			}
		}
	}
	// End Input
	// PrintGrid(sudoku)
	if Solve(sudoku) {
		PrintGrid(sudoku)
	} else {
		fmt.Println("No Solutions!")
	}
}
