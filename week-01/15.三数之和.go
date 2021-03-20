/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := [][]int{}

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}

		if v > 0 {
			return res
		}

		l, r := i+1, len(nums)-1

		for l < r {
			sum := v + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{v, nums[l], nums[r]})
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

// @lc code=end
