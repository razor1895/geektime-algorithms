/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	p, q := 0, 1

	for q < len(nums) {
		if nums[p] != nums[q] {
			if q > p+1 {
				nums[p+1] = nums[q]
			}
			p++
		}
		q++
	}

	return p + 1
}

// @lc code=end

