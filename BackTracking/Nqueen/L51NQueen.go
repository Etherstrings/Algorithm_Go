package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(solveNQueens(4))
}

var res [][]string

func solveNQueens(n int) [][]string {
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}

	backtracking(n, 0, chessboard)

	return res
}

func backtracking(n, level int, chessboard [][]string) {
	if level == n {
		res = append(res, arrayToList(chessboard))
		return
	}

	for i := 0; i < n; i++ {
		if isValid(level, i, n, chessboard) {
			chessboard[level][i] = "Q"
			backtracking(n, level+1, chessboard)
			chessboard[level][i] = "."
		}
	}
}

func isValid(level, col, n int, chessboard [][]string) bool {
	for i := 0; i < level; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}

	for i, j := level-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	for i, j := level-1, col+1; i >= 0 && j <= n-1; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	return true
}

func arrayToList(chessboard [][]string) []string {
	list := make([]string, len(chessboard))
	for i, row := range chessboard {
		list[i] = strings.Join(row, "")
	}
	return list
}
