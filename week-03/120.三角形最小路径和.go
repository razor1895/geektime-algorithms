package main

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
// 从上至下
// func minimumTotal(triangle [][]int) int {
// 	m := len(triangle)
// 	for i := 0; i < m; i++ {
// 		n := len(triangle[i])
// 		for j := 0; j < n; j++ {
// 			if i == 0 && j == 0 {
// 				continue
// 			}
// 			if j == 0 {
// 				triangle[i][j] += triangle[i-1][0]
// 				continue
// 			}
// 			if j == n-1 {
// 				triangle[i][j] += triangle[i-1][j-1]
// 				continue
// 			}
// 			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
// 		}
// 	}

// 	ans := math.MaxInt64

// 	for _, v := range triangle[m-1] {
// 		ans = min(ans, v)
// 	}

// 	return ans
// }
// 从下至上
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j+1], triangle[i+1][j])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dp[i][j] 表示[0,0]到[i,j]的最小路径和
// 状态转移方程
// if j == 0
// dp[i][j] = dp[i-1][j] + triangle[i][j]
// if j == len(triangle[i]) - 1
// dp[i][j] = dp[i-1][j-1] + triangle[i][j]
// else 中间
// dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
// 可以敷用triangle作为dp状态表
// @lc code=end
