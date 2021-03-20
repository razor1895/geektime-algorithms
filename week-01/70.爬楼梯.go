/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	a := 1
	b := 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// @lc code=end

