/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	res, dp_i := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp_i = max(dp_i+nums[i], nums[i])
		res = max(dp_i, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp状态数组 dp[i]表示为到i为止连续子数组的最大和
// dp[i] = max(dp[i-1]+nums[i], nums[i])
// @lc code=end

