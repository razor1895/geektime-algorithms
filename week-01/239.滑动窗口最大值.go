/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	dq := make([]int, 0, k+1)
	res := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		if len(dq) != 0 && i-dq[0] >= k {
			dq = dq[1:]
		}

		for len(dq) != 0 && nums[i] > nums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)

		if i+1 >= k {
			res = append(res, nums[dq[0]])
		}
	}

	return res
}

// @lc code=end
