/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	minCount := len(bank) + 1
	dfs(start, end, bank, make([]bool, len(bank)), 0, &minCount)
	if minCount > len(bank) {
		return -1
	}
	return minCount
}

func dfs(temp string, end string, bank []string, bankVisited []bool, count int, minCount *int) {
	if temp == end {
		if count < *minCount {
			*minCount = count
		}
		return
	}

	for i := 0; i < len(bank); i++ {
		if !bankVisited[i] && checkDiff(temp, bank[i]) {
			bankVisited[i] = true
			dfs(bank[i], end, bank, bankVisited, count+1, minCount)
			bankVisited[i] = false
		}
	}
}

func checkDiff(a string, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	if diff == 1 {
		return true
	}

	return false
}

// @lc code=end
