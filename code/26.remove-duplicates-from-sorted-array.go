package code

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	index := 0
	for _, num := range nums {
		if num != nums[index] {
			index++
			nums[index] = num
		}
	}
	return index + 1
}

// @lc code=end
