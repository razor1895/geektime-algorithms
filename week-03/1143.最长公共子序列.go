/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DP 状态数组定义 dp[i][j]为[0, i-1]的text1数组于[0,j-1]的text2数组
// 的最长公共子序列为dp[i][j]
// 状态转移方程：
// if text1[i] == text[j]
//    dp[i][j] = dp[i-1][j-1] + 1
// else
//    dp[i][j] = max(dp[i-1][j], dp[i][j-1])
// @lc code=end

