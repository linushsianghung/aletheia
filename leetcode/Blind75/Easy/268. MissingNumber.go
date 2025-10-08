package Easy

// https://leetcode.com/problems/missing-number/description
/* Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array. */
func missingNumber(nums []int) int {
	// Just ignoring 0 and using 1 as bottom, so the top is len(nums); and because there is a number missing, so plus 1 to the top
	total := (1 + len(nums)) * len(nums) / 2

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return total - sum
}

func missingNumberExercise(nums []int) int {
	return 0
}
