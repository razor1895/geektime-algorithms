package main

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	currJumpMax := 0
	currPostMax := 0
	jump, i := 0, 0

	for i < len(nums)-1 {
		currPostMax = max(currPostMax, i+nums[i])
		if i == currJumpMax {
			jump++
			currJumpMax = currPostMax
		}
		i++
	}

	return jump
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
