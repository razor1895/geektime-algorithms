/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	k := 0
	for i, v := range nums {
		if i > k {
			return false
		}
		k = max(k, v+i)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

