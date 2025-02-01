package Basic

// https://leetcode.com/problems/rotate-array/description/
/* Given an integer array nums, rotate the array to the right by k steps, where k is non-negative. */
func rotate(nums []int, k int) {
	// step1: reverse the array
	// step2: split the array into two.
	// step3: reverse both the two parts.

	k = k % len(nums)
	swap(nums, 0, len(nums)-1)
	swap(nums, 0, k-1)
	swap(nums, k, len(nums)-1)
}

func swap(nums []int, start int, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
