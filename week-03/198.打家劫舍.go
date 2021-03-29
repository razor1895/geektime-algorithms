/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
// 标准dp
// func rob(nums []int) int {
// 	n := len(nums)
// 	if n == 1 {
// 		return nums[0]
// 	}
// 	dp := make([]int, n)
// 	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

// 	for i := 2; i < n; i++ {
// 		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
// 	}
// 	return dp[n-1]
// }

// 空间压缩, 状态只与dp[i-1]和dp[i-2]有关
func rob(nums []int) int {
	pre, cur := 0, 0
	for _, v := range nums {
		cur, pre = max(pre+v, cur), cur
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp状态定义  dp[i]表示偷到第i间房的最高金额
// 如果偷第i间房，那么i-1不能偷
// dp[i] = dp[i-2] + nums[i]
// 如果不偷第i间房
// dp[i] = dp[i-1]
// 边界条件dp[0]表示一间房子，只能偷，dp[1]表示两间房偷最大的
// @lc code=end

