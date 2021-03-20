package main

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(n2, n3), n5)

		if dp[i] == n2 {
			p2++
		}

		if dp[i] == n3 {
			p3++
		}

		if dp[i] == n5 {
			p5++
		}
	}

	return dp[n-1]
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

// @lc code=end
