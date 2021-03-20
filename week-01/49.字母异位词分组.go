package main

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}

	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	res := make([][]string, 0, len(mp))

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}

// @lc code=end
