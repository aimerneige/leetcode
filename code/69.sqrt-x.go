package code

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x < 0 {
		panic("error")
	}
	if x == 0 || x == 1 {
		return x
	}
	for i := 1; i < x; i++ {
		if (i+1)*(i+1) > x {
			return i
		}
	}

	return x
}

// @lc code=end
