package code

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	hayRunes := []rune(haystack)
	neeLength := len(needle)
	for i := range hayRunes {
		if len(hayRunes[i:]) < neeLength {
			return -1
		}
		if string(hayRunes[i:i+neeLength]) == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
