package main

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	result := fastPow(x, n/2)
	if n%2 == 0 {
		return result * result
	}
	return result * result * x
}

// @lc code=end
