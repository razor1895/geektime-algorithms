/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	count := 0
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j, m, n)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i int, j int, m int, n int) {
	if i < 0 || i >= m || j >= n || j < 0 || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i+1, j, m, n)
	dfs(grid, i-1, j, m, n)
	dfs(grid, i, j+1, m, n)
	dfs(grid, i, j-1, m, n)
}

// @lc code=end
