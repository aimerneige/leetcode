package code

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	i1 := m - 1
	i2 := n - 1
	i := m + n - 1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
		i--
	}
	if i1 >= 0 {
		for ; i >= 0; i-- {
			nums1[i] = nums1[i1]
			i1--
		}
	}
	if i2 >= 0 {
		for ; i >= 0; i-- {
			nums1[i] = nums2[i2]
			i2--
		}
	}

	return
}

// @lc code=end
