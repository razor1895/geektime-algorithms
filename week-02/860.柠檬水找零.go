/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	n5, n10 := 0, 0
	for _, v := range bills {
		if v == 5 {
			n5++
		} else if v == 10 {
			if n5 == 0 {
				return false
			}
			n5--
			n10++
		} else {
			if n10 > 0 && n5 > 0 {
				n10--
				n5--
			} else if n5 > 2 {
				n5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

