package TwoPointers

// TwoSumII https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
/*
Given a 1-indexed array of integer numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.
Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.
Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.
*/
func TwoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

func twoSumIIExercise(numbers []int, target int) []int {
	return nil
}

/*
# Analysis:
1. Concept: The "Two Pointers" technique. Because the array is sorted, we have a property of monotonicity.
   Moving the left pointer right always increases (or maintains) the sum, and moving the right pointer left
   always decreases (or maintains) the sum. This allows us to converge on the target in a single pass.

2. How to derive the optimized solution during an Interview:
   - Step 1: Start with Brute Force. Mention that checking every pair takes O(n^2).
   - Step 2: Leverage the "Sorted" property. Ask yourself: "If I pick a number, can I find its complement faster?"
     This leads to Binary Search for each element, resulting in O(n log n).
   - Step 3: Optimize further to O(n). Realize that if the sum of numbers[left] + numbers[right] is too small,
     incrementing 'left' is the only way to potentially reach the target. Conversely, if it's too large,
     decrementing 'right' is the only way. This eliminates the need for redundant checks.

3. Complexity Analysis:
   - Time Complexity: O(n), where n is the length of the array. Each element is visited at most once.
   - Space Complexity: O(1), as we only store two integer pointers regardless of input size.

4. Interview Suggestion: Always clarify if the array can contain duplicates (it doesn't change the logic here
   but shows attention to detail) and confirm if the result should be 0-indexed or 1-indexed.

Implementation Method:
public int[] twoSum(int[] numbers, int target) {
    int left = 0;
    int right = numbers.length - 1;

    while (left < right) {
        int sum = numbers[left] + numbers[right];

        if (sum == target) {
            return new int[]{left + 1, right + 1}; // 1-indexed requirement
        } else if (sum < target) {
            left++; // Need a larger sum
        } else {
            right--; // Need a smaller sum
        }
    }
    return new int[]{-1, -1}; // Should not be reached per constraints
}
*/
