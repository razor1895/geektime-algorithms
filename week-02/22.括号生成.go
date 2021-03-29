package main

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}

	backtrack(&res, "", n, n, n)
	return res
}

func backtrack(res *[]string, pairs string, l int, r int, n int) {
	if l == 0 && r == 0 {
		*res = append(*res, pairs)
		return
	}
	if l > r {
		return
	}
	if l > 0 {
		backtrack(res, pairs+"(", l-1, r, n)
	}
	if r > 0 {
		backtrack(res, pairs+")", l, r-1, n)
	}
}

// @lc code=end
