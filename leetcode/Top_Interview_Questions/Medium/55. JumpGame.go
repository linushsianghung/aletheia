package Medium

// CanJump https://leetcode.com/problems/jump-g	ame/
/*
You are given an integer array nums. You are initially positioned at the array's first index,
and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/
func CanJump(nums []int) bool {
	anchor := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		// if nums[i] is larger than the gap between current position i and anchor, it means the anchor position can be reached.
		// Then we move anchor backward to position i and check if this new anchor position can be reached as well
		if nums[i]+i >= anchor {
			anchor = i
		}
	}

	return anchor == 0
}

func canJumpExercise(nums []int) bool {

	return false
}
