package code

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	index := 0
	for _, num := range nums {
		if num != val {
			nums[index] = num
			index++
		}
	}
	return index
}

// @lc code=end
