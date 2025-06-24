# Binary Search
  Ref:
- [Binary Search in Detail](https://labuladong.gitbook.io/algo-en/iii.-algorithmic-thinking/detailedbinarysearch)
- [Binary Search 101](https://leetcode.com/problems/binary-search/solutions/423162/binary-search-101/)
- [Binary Search: A Comprehensive Guide](https://leetcode.com/discuss/study-guide/3726061/binary-search-a-comprehensive-guide)
- [Are you one of the 10% of programmers who can write a binary search?](https://reprog.wordpress.com/2010/04/19/are-you-one-of-the-10-percent/)

## Concepts
Binary Search operates on a `sorted sequence with a specified left and right index`. This is called the **Search Space**. Binary Search maintains the left, right,
and middle indices of the search space and compares the *search target* or *applies the search condition* to the middle value of the collection; if the condition is
unsatisfied or values unequal, the half in which the target cannot lie is eliminated and the search continues on the remaining half until it is successful.
If the search ends with an empty half, the condition cannot be fulfilled and the target is not found.

### When can we use Binary Search:
- When need to search for an index or element in a space and if the space is unordered, we can always sort it first before applying Binary Search.
> Situation: Array => Sorted Array => Binary Search (or more general, 2 pointers)
- If we can discover some kind of **monotonicity**, for example, if condition(k) is True then condition(k + 1) is True (*implicitly sorted space*), then we can consider Binary Search .
- **Binary search sorted assumption is necessary only when there is only 1 "unique value" in the array. (Refer to [162. Find Peak Element](../../../leetcode/LeetCode75/BinarySearch/162.%20FindPeakElement.go))** 

### Template
```Golang
package BinarySearch

import "github.com/linushung/aletheia/Concepts/algorithm/BinarySearch"

/*
704. Binary Search: https://leetcode.com/problems/binary-search/description/

Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
If target exists, then return its index. Otherwise, return -1.

You must write an Algorithm with O(log n) runtime complexity.
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// Calculate mid := left + (right-left)/2 rather than (right-left)/2 to prevent overflow.
		// Here we decide using lower mid rather than upper mid (left + (right-left+1)/2)
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
```
if you are using lower mid mid = low + (high - low)/2 (the mid is considered as low side), then you NEED to use high = mid & low = mid + 1 to shrink the boundary. 
And if you are using upper mid mid = low + (high - low + 1)/2 (the mid is considered as high side) , you NEED to use low = mid & high = mid -1 to shrink boundary.
## Examples
```Golang
package BinarySearch

import "github.com/linushung/aletheia/Concepts/algorithm/BinarySearch"

/*
35. Search Insert Position: https://leetcode.com/problems/search-insert-position/description/

Given a sorted array of distinct integers and a target value, return the index if the target is found.If not, return the index where it would be if it were inserted in order.

You must write an Algorithm with O(log n) runtime complexity.
*/
func searchInsert(nums []int, target int) int {
	return BinarySearch.SearchInsert(nums, target)
}
/*
69. Sqrt(x): https://leetcode.com/problems/sqrtx/description/

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.
*/
func mySqrt(x int) int {
	return BinarySearch.MySqrt(x)
}

/*
875. Koko Eating Bananas: https://leetcode.com/problems/koko-eating-bananas/description/

Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.
*/
func minEatingSpeed(piles []int, h int) int {
	return BinarySearch.MinEatingSpeed(piles, h)
}
```
