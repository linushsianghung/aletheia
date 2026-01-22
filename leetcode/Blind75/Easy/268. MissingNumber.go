package Easy

// https://leetcode.com/problems/missing-number/description
/* Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array. */
func missingNumber(nums []int) int {
	// The largest numbe is supposed to be len(nums)-1 (e.g. 0 ~ 10 should be (11 - 1)). Because is a 0 based array, we can just use len(array) instead.
	total := (0 + len(nums)) * (len(nums) + 1) / 2

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return total - sum
}

func missingNumberExercise(nums []int) int {
	return 0
}
