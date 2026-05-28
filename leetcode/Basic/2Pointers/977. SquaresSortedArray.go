package TwoPointers

// https://leetcode.com/problems/squares-of-a-sorted-array/
/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
*/
func sortedSquares(nums []int) []int {
	left, right, tail := 0, len(nums)-1, len(nums)-1
	result := make([]int, len(nums))

	for left <= right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			result[tail] = nums[left] * nums[left]
			left++
		} else {
			result[tail] = nums[right] * nums[right]
			right--
		}
		tail--
	}
	return result
}

func sortedSquaresExercise(nums []int) []int {
	return nil
}

/*
# Analysis:
1. Concept:
   The input array is sorted in non-decreasing order. However, squaring negative numbers makes them positive,
   potentially larger than the squares of positive numbers. This means the largest squared values will always
   reside at the boundaries (the far left or the far right) of the input array.
   By using two pointers (left and right), we can compare the squares at both ends and fill a result array
   from the largest value to the smallest (tail to head).

2. Complexity Analysis:
   - Time Complexity: O(n), where n is the length of the array. We iterate through the array once using
     the two-pointer approach.
   - Space Complexity: O(n) to store the result array. The auxiliary space used (pointers and temp variables) is O(1).

3. Interview Suggestion:
   - Step 1: Clarify if the input can contain negative numbers. If not, a simple one-pass squaring is enough.
   - Step 2: Mention the Brute Force (O(n log n)): Square all elements and then sort.
   - Step 3: Propose the Optimization (O(n)): Explain that the "magnitude" of elements is highest at the ends.
   - Step 4: Implement using two pointers and fill the result array backwards to maintain the non-decreasing order.

Implementation Method:
public int[] sortedSquares(int[] nums) {
	int[] result = new int[nums.length];
	int left = 0, right = nums.length-1, tail = nums.length-1;

	while (left <= right) {
		int squareLeft = nums[left]*nums[left];
		int squareRight = nums[right]*nums[right];

		if (squareLeft > squareRight) {
			result[tail] = squareLeft;
			left++;
		} else {
			result[tail] = squareRight;
			right--;
		}
		tail--;
	}

	return result;
}
*/
