/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n < 2 {
		return 0
	}
	return max(robHelper(nums[1:]), robHelper(nums[:n-1]))
}

func robHelper(nums []int) int {
	pre, cur := 0, 0
	for _, v := range nums {
		cur, pre = max(v+pre, cur), cur
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp状态定义和状态转移方程同198题，但是因为这里的屋子不能同时偷最后一间和第一间
// 这里可以分成三种情况，偷第一间不偷最后一间，第一和最后都不偷和偷最后一间
// @lc code=end

