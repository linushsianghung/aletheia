package Array_String

import "math"

// https://leetcode.com/problems/increasing-triplet-subsequence/description/?envId=leetcode-75
/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k].
If no such indices exists, return false.

Analysis:
- Reference: https://leetcode.com/problems/increasing-triplet-subsequence/solutions/78993/clean-and-short-with-comments-c/comments/191930/
Triple2 is only set when there is a triple1 before it which is smaller. It's impossible for triple2 to be set to a certain number and not have a number before it be smaller.
So there must be a subsequence, even though {triple1, triple2, triple3} is not necessarily the correct subsequence.

- Reference: https://leetcode.com/problems/increasing-triplet-subsequence/solutions/78993/clean-and-short-with-comments-c/comments/83820/
The solution seems to fail for the input where the smallest element's index surpasses the second-smallest one's, but actually it doesn't matter that triple1 is the smallest index or not.
Consider [2, 6, 1, 8] as example. triple1 here would get updated to 1, triple2 to 6 and the 3rd element is 8.
Even though second largest and largest elements are correct, the correct order of subsequence should be 2, 6, 8 and not 6, 1, 8.
Thus, if the problem requires us to return the index, then this code would not work.
*/
// Reference: https://leetcode.com/problems/increasing-triplet-subsequence/solutions/78993/clean-and-short-with-comments-c
func increasingTriplet(nums []int) bool {
	triplet1, triplet2 := math.MaxInt, math.MaxInt

	for _, num := range nums {
		// It is necessary to be "less and equal to", in case all the numbers are the same
		if num <= triplet1 {
			triplet1 = num // triple1 is so far the smallest number which is a candidate for 1st element
		} else if num <= triplet2 {
			triplet2 = num // There is always a number less than triplet2, but might not be triplet1 (otherwise, triple2 == math.MaxInt and won't get updated)
		} else {
			// At this point, we have/had triplet1 < triplet2 already and num > triplet2
			return true
		}
	}

	return false
}

func increasingTripletExercise(nums []int) bool {
	return false
}
