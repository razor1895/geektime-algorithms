package main

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
	return max(robHouse(nums[1:]), robHouse(nums[:n-1]))
}

func robHouse(nums []int) int {
	cur, pre := 0, 0
	for _, num := range nums {
		cur, pre = max(pre+num, cur), cur
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
