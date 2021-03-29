/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := [][]string{}
	backtrack(n, &res, make([]int, 0, n), 0)
	return res
}

func backtrack(n int, res *[][]string, board []int, row int) {
	if row == n {
		temp := make([]string, n)
		for i, v := range board {
			temp[i] = strings.Repeat(".", v) + "Q" + strings.Repeat(".", n-1-v)
		}
		*res = append(*res, temp)
		return
	}

	for col := 0; col < n; col++ {
		if !checkValid(board, col, row) {
			board = append(board, col)
			backtrack(n, res, board, row+1)
			board = board[:len(board)-1]
		}
	}
}

func checkValid(board []int, col int, row int) bool {
	for itemRow, itemCol := range board {
		if itemCol == col || itemRow+itemCol == col+row || itemCol-itemRow == col-row {
			return true
		}
	}

	return false
}

// @lc code=end
