/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	cnt := amount + 1
	dp := make([]int, cnt)

	for v := range dp {
		dp[v] = cnt
	}

	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == cnt {
		return -1
	}

	return dp[amount]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dp状态定义 dp[i]表示金额为i时最少的硬币个数
// dp[i] = min(dp[i], dp[i-conins[j]] + 1)
// @lc code=end

