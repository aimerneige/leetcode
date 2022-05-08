package code

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	li := 0             // left index
	ri := len(nums) - 1 // right index
	for li != ri {
		mid := (ri-li)/2 + li
		if li+1 == ri {
			if target == nums[li] {
				return li
			}
			if target == nums[ri] {
				return ri
			}
			if target > nums[ri] {
				return ri + 1
			}
			if target < nums[li] {
				return ri - 1
			}
			return ri
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			li = mid
		} else {
			ri = mid
		}
	}
	if nums[li] >= target {
		return li
	}
	return li + 1
}

// @lc code=end
