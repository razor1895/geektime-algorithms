/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// 一次遍历所有数字出现的频次, { 1: 3, 2: 2, 3: 1 }
	mapper := make(map[int]int)
	for _, v := range nums {
		mapper[v]++
	}
	//桶排序，构建二维数组，O(n),
	stack := make([][]int, len(nums))
	for k, freq := range mapper {
		stack[freq-1] = append(stack[freq-1], k)
	}
	//O(k)
	var res []int
	for i := len(stack) - 1; i >= 0; i-- {
		if len(stack[i]) == 0 {
			continue
		}
		res = append(res, stack[i]...)
		if len(res) >= k {
			break
		}
	}
	return res
}

// @lc code=end

