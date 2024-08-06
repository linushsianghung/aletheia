package Medium

// https://leetcode.com/problems/jump-game/
/*
You are given an integer array nums. You are initially positioned at the array's first index, 
and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/
func canJump(nums []int) bool {
	tail := len(nums)-1

	for i := tail-1; i >= 0; i-- {
		if i + nums[i] >= tail {
			tail = i
		}
	}

	return tail == 0
}
