/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n < 3 {
		return n
	}
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// dp table定义 dp[i] 为从0到i一共有多少种爬法
// 状态转移方程 dp[i] = dp[i-1] + dp[i-2]
// @lc code=end

