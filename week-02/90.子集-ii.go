/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	backTrack(nums, &res, []int{}, 0)
	return res
}

func backTrack(nums []int, res *[][]int, temp []int, start int) {

	track := make([]int, len(temp))
	copy(track, temp)
	*res = append(*res, track)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i-1] == nums[i] {
			continue
		}
		temp = append(temp, nums[i])
		backTrack(nums, res, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}

// @lc code=end
