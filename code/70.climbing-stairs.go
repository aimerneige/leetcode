package code

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	result := 2
	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		result += min(a, b)
		if a < b {
			a = result
		} else {
			b = result
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
