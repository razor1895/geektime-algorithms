/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	size := len(heights)
	stack := make([]int, 0, size-1)
	stack = append(stack, -1)
	res := 0
	i := 0

	for i < size && len(stack) != 0 {
		cur := heights[i]
		topIndex := stack[len(stack)-1]
		if topIndex < 0 || cur >= heights[topIndex] {
			stack = append(stack, i)
			i++
			continue
		}

		stack = stack[:len(stack)-1]
		begin := stack[len(stack)-1]
		area := (i - begin - 1) * heights[topIndex]
		res = max(res, area)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end
