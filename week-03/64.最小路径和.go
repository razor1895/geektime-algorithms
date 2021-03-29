package main

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j > 0 {
					grid[0][j] += grid[0][j-1]
				}
				continue
			}
			if j == 0 {
				if i > 0 {
					grid[i][0] += grid[i-1][0]
				}
				continue
			}
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dp[i][j] 表示从[0, 0]到[i, j]这个点的最小路径和
// 状态转移方程
// dp[i][j] = min(dp[i-1][j], dp[i], j-1) + grid[i][j]
// @lc code=end
