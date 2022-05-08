package code

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := make([]rune, 10000)
	index := -1
	isFull := func() bool {
		return index == 9999
	}
	isEmpty := func() bool {
		return index == -1
	}
	push := func(r rune) bool {
		if !isFull() {
			index++
			stack[index] = r
			return true
		}
		return false
	}
	pop := func() (rune, bool) {
		if !isEmpty() {
			v := stack[index]
			index--
			return v, true
		}
		return ' ', false
	}
	runes := []rune(s)
	for _, r := range runes {
		switch r {
		case '(':
			if !push(r) {
				return false
			}
		case ')':
			v, ok := pop()
			if !ok {
				return false
			}
			if v != '(' {
				return false
			}
		case '[':
			if !push(r) {
				return false
			}
		case ']':
			v, ok := pop()
			if !ok {
				return false
			}
			if v != '[' {
				return false
			}
		case '{':
			if !push(r) {
				return false
			}
		case '}':
			v, ok := pop()
			if !ok {
				return false
			}
			if v != '{' {
				return false
			}
		}
	}
	return isEmpty()
}

// @lc code=end
