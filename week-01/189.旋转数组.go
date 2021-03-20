/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverseArray(nums, 0, n-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, n-1)
}
func reverseArray(nums []int, start int, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start, end = start+1, end-1
	}
}

// @lc code=end

