package main

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	res := [][]string{}
	backtrack(s, &res, &[]string{}, 0)
	return res
}

func backtrack(s string, res *[][]string, temp *[]string, start int) {
	if start == len(s) {
		path := make([]string, len(*temp))
		copy(path, *temp)
		*res = append(*res, path)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) == false {
			continue
		}
		*temp = append(*temp, s[start:i+1])
		backtrack(s, res, temp, i+1)
		*temp = (*temp)[:len(*temp)-1]
	}
}

func isPalindrome(s string, l int, r int) bool {
	if l == r {
		return true
	}
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// @lc code=end
