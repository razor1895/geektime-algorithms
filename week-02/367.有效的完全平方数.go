package main

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	l, r := 2, num/2

	for l <= r {
		x := l + (r-l)/2
		mid := x * x
		if mid == num {
			return true
		}

		if mid > num {
			r = r - 1
		} else {
			l = l + 1
		}
	}

	return false
}

// @lc code=end
