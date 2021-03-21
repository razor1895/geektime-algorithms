package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
var res4 [][]int

// 给定一个元素
func subsets(nums []int) [][]int {
	res4 = [][]int{}
	backTrack1([]int{}, nums, 0)
	return res4
}

func backTrack1(temp []int, nums []int, start int) {

	track := make([]int, len(temp))
	copy(track, temp)
	res4 = append(res4, track)

	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backTrack1(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

// @lc code=end
