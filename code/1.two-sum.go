package code

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		require := target - num
		if v, ok := indexMap[require]; ok {
			return []int{i, v}
		}
		indexMap[num] = i
	}
	return []int{}
}

// @lc code=end
