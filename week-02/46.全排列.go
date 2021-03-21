package main

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, &res, []int{})
	return res
}

func backtrack(nums []int, res *[][]int, temp []int) {
	if len(temp) == len(nums) {
		path := make([]int, len(temp))
		copy(path, temp)
		*res = append(*res, path)
	}

	for i := 0; i < len(nums); i++ {
		if contains(temp, nums[i]) {
			continue
		}
		temp = append(temp, nums[i])
		backtrack(nums, res, temp)
		temp = temp[:len(temp)-1]
	}
}

func contains(nums []int, target int) bool {
	for _, v := range nums {
		if target == v {
			return true
		}
	}

	return false
}

// @lc code=end
