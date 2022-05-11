package code

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	length := 0
	runes := []rune(s)
	find := false
	for i := 0; i < len(s); i++ {
		if runes[i] == ' ' && !find {
			continue
		}
		if runes[i] != ' ' && !find {
			find = true
			length = 0
			length++
			continue
		}
		if runes[i] == ' ' && find {
			find = false
			continue
		}
		if runes[i] != ' ' && find {
			length++
			continue
		}
	}

	return length
}

// @lc code=end
