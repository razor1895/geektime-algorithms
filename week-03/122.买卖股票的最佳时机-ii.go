/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	dp_i_0, dp_i_1 := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		dp_i_0, dp_i_1 = max(dp_i_1+prices[i], dp_i_0), max(dp_i_0-prices[i], dp_i_1)
	}

	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

