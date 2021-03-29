package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	DIRECTION := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 北东南西
	direction, x, y, ans := 0, 0, 0, 0
	maps := map[string]bool{}
	for _, v := range obstacles {
		maps[strconv.Itoa(v[0])+","+strconv.Itoa(v[1])] = true
	}
	for _, com := range commands {
		if com >= 0 {
			new_x, new_y := 0, 0
			for j := 0; j < com; j++ {
				new_x, new_y = x+DIRECTION[direction][0], y+DIRECTION[direction][1]
				if _, ok := maps[strconv.Itoa(new_x)+","+strconv.Itoa(new_y)]; ok {
					break
				}
				x, y = new_x, new_y
				ans = max(ans, x*x+y*y)
			}
		} else if com == -1 {
			direction = (direction + 1) % 4
		} else {
			direction = (direction + 3) % 4
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
