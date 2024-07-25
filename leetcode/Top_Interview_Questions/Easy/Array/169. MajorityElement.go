package Array

import "sort"

// https://leetcode.com/problems/majority-element/description/
// Ref:
// - https://leetcode.com/problems/majority-element/solutions/3676530/3-method-s-beats-100-c-java-python-beginner-friendly/?envType=featured-list&envId=top-interview-questions?envType=featured-list&envId=top-interview-questions
// - NeetCode: https://www.youtube.com/watch?v=7pnhv842keE
/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
*/
func majorityElement(nums []int) int {
	return majorityElementMooreVoting(nums)
	// return majorityElementHashMap(nums)
	// return majorityElementSort(nums)
}

/*
Algorithm:
The intuition behind the Moore's Voting Algorithm is based on the fact that if there is a majority element in an array,
it will always remain in the lead, even after encountering other elements.
*/
func majorityElementMooreVoting(nums []int) int {
	result, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			result = num
		}

		if result == num {
			count++
		} else {
			count--
		}
	}

	return result
}

/*
Algorithm:
The intuition behind using a hash map is to count the occurrences of each element in the array and then identify the element that occurs more than n/2 times.
*/
func majorityElementHashMap(nums []int) int {
	note := make(map[int]int)

	for _, num := range nums {
		note[num]++

		if note[num] > len(nums)/2 {
			return num
		}
	}

	return 0
}

/*
Algorithm:
The intuition behind this approach is that if an element occurs more than n/2 times in the array (where n is the size of the array),
it will always occupy the middle position when the array is sorted.
*/
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
