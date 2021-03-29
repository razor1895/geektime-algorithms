/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	dp_i_0, dp_i_1 := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		dp_i_0 = max(dp_i_0, prices[i]+dp_i_1)
		dp_i_1 = max(dp_i_1, -prices[i])
	}

	return dp_i_0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 设计状态数组 dp[i][0], dp[i][1] 表示第i天持有股票第收入
// dp[i][0]表示没有持有股票第收入，包括没有买入或者卖出了
// dp[i][1] = max(dp_i_1, - nums[i])
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + nums[i])
// @lc code=end
