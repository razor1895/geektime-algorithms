/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums, &res, []int{}, &used)
	return res
}

func backtrack(nums []int, res *[][]int, temp []int, used *[]bool) {
	if len(temp) == len(nums) {
		path := make([]int, len(temp))
		copy(path, temp)
		*res = append(*res, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[i] || i > 0 && nums[i-1] == nums[i] && !(*used)[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		(*used)[i] = true
		backtrack(nums, res, temp, used)
		(*used)[i] = false
		temp = temp[:len(temp)-1]
	}
}

// @lc code=end
