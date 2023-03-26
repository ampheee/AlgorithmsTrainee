package leetcode

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:
//
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.
//func isValidSudoku(board [][]byte) bool {
//	var res = true
//	rows := make([]map[byte]bool, 9)
//	cols := make([]map[byte]bool, 9)
//	boxes := make([]map[byte]bool, 9)
//	for i := 0; i < 9; i++ {
//		rows[i] = make(map[byte]bool)
//		cols[i] = make(map[byte]bool)
//		boxes[i] = make(map[byte]bool)
//	}
//	for i := 0; i < 9; i++ {
//		for j := 0; j < 9; j++ {
//			if board[i][j] != '.' {
//				num := board[i][j]
//				boxIndex := (i/3)*3 + j/3
//				if rows[i][num] || cols[j][num] || boxes[boxIndex][num] {
//					return false
//				}
//				rows[i][num] = true
//				cols[j][num] = true
//				boxes[boxIndex][num] = true
//			}
//		}
//	}
//	return res
//}
