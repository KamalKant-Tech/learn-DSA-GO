package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(solveNQueens(4))
	// fmt.Println(totalNQueens(4))

	//create a sudoku matrix of 9*9 elements
	//matrix := make([][]byte, 9)
	matrix := [][]byte{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0 ,0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	solveSudoku(matrix)
	// fmt.Println(isValidSudoku(matrix))

	// for row := 0; row < len(matrix); row++ {
	// 	for col := 0; col < len(matrix[0]); col++ {
	// 		matrix[row] = append(matrix[row], )
	// 	}
	// }
}

/**
* Problem: 51. N-Queens
* The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
* Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
* Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
* Input: n = 4
* Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
* Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
 */

func solveNQueens(n int) [][]string {
	// Create n*n matrix
	matrix := make([][]rune, n)
	ans := [][]string{}
	for i := 0; i < n; i++ {
		// str := []string{}
		matrix[i] = make([]rune, n)

		for j := 0; j < n; j++ {
			// str = append(str, ".")
			matrix[i][j] = '.'
		}

		// matrix[i] = append(matrix[i], str...)
	}
	leftRow := make([]int, n)
	upperdiagonal := make([]int, 2*n-1)
	loweDiagonal := make([]int, 2*n-1)
	fmt.Println(loweDiagonal)
	// solveQueens(&matrix, 0, n, &ans)
	solveQueensOptimal(0, &matrix, &ans, n, &leftRow, &upperdiagonal, &loweDiagonal)
	return ans
}

func solveQueens(matrix *[][]rune, col int, n int, ans *[][]string) {
	if col == n {
		addToResult(*matrix, ans)
		return
	}

	for row := 0; row < n; row++ {
		if isSafe(row, col, *matrix, n) {
			(*matrix)[row][col] = 'Q'
			solveQueens(matrix, col+1, n, ans)
			(*matrix)[row][col] = '.'
		}
	}

}

// This func checks the queen in upperleftdiagonal, left and lowwerleftdiagonal

func isSafe(row, col int, matrix [][]rune, n int) bool {
	dupRow, dupCol := row, col

	// This is for upperdiagonal
	for row >= 0 && col >= 0 {
		if matrix[row][col] == 'Q' {
			return false
		}
		row--
		col--
	}

	row, col = dupRow, dupCol
	// This is for left side column
	for col >= 0 {
		if matrix[row][col] == 'Q' {
			return false
		}
		col--
	}

	row, col = dupRow, dupCol
	// This is for left side column
	for row < n && col >= 0 {
		if matrix[row][col] == 'Q' {
			return false
		}
		row++
		col--
	}
	return true
}

func addToResult(board [][]rune, res *[][]string) {
	n := len(board)
	rows := []string{}

	for r := 0; r < n; r++ {
		row := strings.Builder{}
		for c := 0; c < n; c++ {
			row.WriteRune(board[r][c])
		}
		rows = append(rows, row.String())
	}

	*res = append(*res, rows)
}

func solveQueensOptimal(col int, matrix *[][]rune, ans *[][]string, n int, leftRow *[]int, uppperDiagonal *[]int, lowerDiagonal *[]int) {
	if col == n {
		addToResult(*matrix, ans)
		return
	}

	for row := 0; row < n; row++ {
		if (*leftRow)[row] == 0 && (*uppperDiagonal)[n-1+col-row] == 0 && (*lowerDiagonal)[row+col] == 0 {
			(*matrix)[row][col] = 'Q'
			(*leftRow)[row] = 1
			(*uppperDiagonal)[n-1+col-row] = 1
			(*lowerDiagonal)[row+col] = 1
			solveQueensOptimal(col+1, matrix, ans, n, leftRow, uppperDiagonal, lowerDiagonal)
			(*matrix)[row][col] = '.'
			(*leftRow)[row] = 0
			(*uppperDiagonal)[n-1+col-row] = 0
			(*lowerDiagonal)[row+col] = 0
		}
	}
}

/**
* Problem: 52. N-Queens II
* The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
* Given an integer n, return the number of distinct solutions to the n-queens puzzle.
* Input: n = 4
* Output: 2
* Explanation: There are two distinct solutions to the 4-queens puzzle as shown.
* Approach: Going to use same approach as above will n*n matrix along with this left, upperdigonal and lowerdiagonal rows to keep track of queens position.
* Here instead of using ans slice of string, I can use the ans as int type and pass it by reference in the function.
*/

func totalNQueens(n int) int {
    // Create n*n matrix
	matrix := make([][]rune, n)
    var ans int
	for i := 0; i < n; i++ {
		matrix[i] = make([]rune, n)

		for j := 0; j < n; j++ {
			matrix[i][j] = '.'
		}
	}
	leftRow := make([]int, n)
	upperdiagonal := make([]int, 2*n-1)
	loweDiagonal := make([]int, 2*n-1)
	solveQueensTotalOptimal(0, &matrix, &ans, n, &leftRow, &upperdiagonal, &loweDiagonal)
	return ans
}

func solveQueensTotalOptimal(col int, matrix *[][]rune, ans *int, n int, leftRow *[]int, uppperDiagonal *[]int, lowerDiagonal *[]int) {
	if col == n {
        *ans++
		return
	}

	for row := 0; row < n; row++ {
		if (*leftRow)[row] == 0 && (*uppperDiagonal)[n-1+col-row] == 0 && (*lowerDiagonal)[row+col] == 0 {
			(*matrix)[row][col] = 'Q'
			(*leftRow)[row] = 1
			(*uppperDiagonal)[n-1+col-row] = 1
			(*lowerDiagonal)[row+col] = 1
			solveQueensTotalOptimal(col+1, matrix, ans, n, leftRow, uppperDiagonal, lowerDiagonal)
			(*matrix)[row][col] = '.'
			(*leftRow)[row] = 0
			(*uppperDiagonal)[n-1+col-row] = 0
			(*lowerDiagonal)[row+col] = 0
		}
	}
}

/**
* Problem: 37. Sudoku Solver
*/

func solveSudoku(board [][]byte)  {
	solve(&board)
	fmt.Println(board)
}

func solve(board *[][]byte) bool {
	for row := 0; row < len(*board); row++ {
		for col := 0; col < len((*board)[0]); col++ {
			if (*board)[row][col] == '.' {
				for val := byte('1'); val <= byte('9'); val++ {
					if isValid(board, row, col, val) {
						(*board)[row][col] = byte(val)
						if solve(board) {
							return true
						} else {
							(*board)[row][col] = '.'
						}
					}
				}
				return false
			}
		}
	}

	return true
}
func isValid(board *[][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if (*board)[i][col] == val {
			return false
		}

		if (*board)[row][i] == val {
			return false
		}

		iRow := (3 * (row/3)) + (i / 3)
		iCol := (3 * (col/3)) + (i % 3)
		fmt.Println("Is this failing here", iRow, iCol)

		if (*board)[iRow][iCol] == val {
			return false
		}
	}
	return true
}
