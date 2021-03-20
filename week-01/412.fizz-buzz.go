/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	ans := []string{}
	for i := 1; i <= n; i++ {
		d3, d5 := i%3 == 0, i%5 == 0
		if d3 && d5 {
			ans = append(ans, "FizzBuzz")
		} else if d3 {
			ans = append(ans, "Fizz")
		} else if d5 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}

// @lc code=end

