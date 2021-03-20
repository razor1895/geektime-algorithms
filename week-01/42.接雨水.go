package main

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	stack := make([]int, len(height))
	ans := 0

	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			l, r := stack[len(stack)-1], i
			h := min(height[l], height[r]) - height[cur]
			ans += (r - l - 1) * h
		}
		stack = append(stack, i)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
