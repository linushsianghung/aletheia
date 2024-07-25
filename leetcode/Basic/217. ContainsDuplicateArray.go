package Basic

// https://leetcode.com/problems/contains-duplicate
/* Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct. */
func containsDuplicate(nums []int) bool {
	note := make(map[int]bool)

	for _, num := range nums {
		if _, ok := note[num]; ok {
			return true
		}
		note[num] = true
	}

	return false
}

func containsDuplicateExercise(nums []int) bool {

	return false
}
