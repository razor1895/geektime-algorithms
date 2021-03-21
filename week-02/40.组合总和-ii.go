package main

import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	backtrack(candidates, &res, &[]int{}, target, 0)
	return res
}

func backtrack(nums []int, res *[][]int, temp *[]int, sum int, start int) {
	if sum < 0 {
		return
	}
	if sum == 0 {
		path := make([]int, len(*temp))
		copy(path, *temp)
		*res = append(*res, path)
		return
	}

	for i := start; i < len(nums); i++ {
		if i > start && nums[i-1] == nums[i] {
			continue
		}
		*temp = append(*temp, nums[i])
		backtrack(nums, res, temp, sum-nums[i], i+1)
		*temp = (*temp)[:len(*temp)-1]
	}
}

// @lc code=end
