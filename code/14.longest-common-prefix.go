package code

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	result := ""
	runes := []rune(strs[0])
	for i, r := range runes {
		for _, s := range strs {
			if len([]rune(s)) <= i {
				return result
			}
			if []rune(s)[i] != r {
				return result
			}
		}
		result += string(r)
	}
	return result
}

// @lc code=end
