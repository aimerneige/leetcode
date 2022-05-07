package code

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	num := 0
	for num < x {
		num *= 10
		t := x % 10
		num += t
		if num == x || num == x/10 {
			return true
		}
		x /= 10
	}
	return false
}

// @lc code=end
