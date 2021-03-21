/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	major, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != major {
			count--
		} else {
			count++
		}

		if count == 0 {
			major = nums[i]
			count = 1
		}
	}

	return major
}

// @lc code=end

