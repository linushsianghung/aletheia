# Binary Search
Ref:
- LeetCode Explore: https://leetcode.com/explore/learn/card/binary-search/
- NeetCode: https://www.youtube.com/playlist?list=PLot-Xpze53leNZQd0iINpD-MAhMOMzWvO
- mycodeschool: https://www.youtube.com/playlist?list=PL2_aWCzGMAwL3ldWlrii6YeLszojgH77j

## Concepts
Binary Search operates on a contiguous sequence with a specified left and right index. This is called the Search Space. Binary Search maintains the left, right,
and middle indices of the search space and compares the search target or applies the search condition to the middle value of the collection; if the condition is
unsatisfied or values unequal, the half in which the target cannot lie is eliminated and the search continues on the remaining half until it is successful.
If the search ends with an empty half, the condition cannot be fulfilled and the target is not found.

Binary Search should be considered every time you need to search for an index or element in a collection. If the collection is unordered, we can always sort it first before applying Binary Search.
Situation: Array => Sorted Array => Binary Search (or more generally, 2 pointers)

## Examples
```Golang
package BinarySearch

import "github.com/linushung/aletheia/leetcode/Top_Interview_Questions/Medium"
/*
704. Binary Search: https://leetcode.com/problems/binary-search/description/

Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/
func search(nums []int, target int) int {
	return Search(nums, target)
}

/*
34. Find First and Last Position of Element in Sorted Array: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.
*/
func searchRange(nums []int, target int) []int {
	return Medium.SearchRange(nums, target)
}

/*
2300. Successful Pairs of Spells and Potions (https://leetcode.com/problems/successful-pairs-of-spells-and-potions/description/)

You are given two positive integer arrays spells and potions, of length n and m respectively, where spells[i] represents the strength of the ith spell and
potions[j] represents the strength of the jth potion.
You are also given an integer success. A spell and potion pair is considered successful if the product of their strengths is at least success.
Return an integer array pairs of length n where pairs[i] is the number of potions that will form a successful pair with the ith spell.

Analysis:
1. Brute Force: Double for loop for spell and potion arrays => Not Accepted
2. Improvement
- Only need to consider "successful" cases, which means that just need to find out where the product of the strengths can be larger than "success"
- If potions are sorted, we can use binary search to find out the target product
*/

func successfulPairs(spells []int, potions []int, success int64) []int {
	return nil
}
```
