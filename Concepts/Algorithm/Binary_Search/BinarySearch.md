# Binary Search
  Ref:
- [Binary Search in Detail](https://labuladong.gitbook.io/algo-en/iii.-algorithmic-thinking/detailedbinarysearch)
- [Binary Search: A Comprehensive Guide](https://leetcode.com/discuss/study-guide/3726061/binary-search-a-comprehensive-guide)
- [Are you one of the 10% of programmers who can write a binary search?](https://reprog.wordpress.com/2010/04/19/are-you-one-of-the-10-percent/)

## Concepts
### Basic Idea
Binary Search operates on a `sorted sequence with a specified left and right index`. This is called the **Search Space**. Binary Search maintains the left, right, and middle indices of the search space and compares the *search target* or *applies the search condition* to the middle value of the collection; if the condition is unsatisfied or values unequal, the half in which the target cannot lie is eliminated and the search continues on the remaining half until it is successful. If the search ends with an empty half, the condition cannot be fulfilled and the target is not found.

### The Core Principle: Monotonicity***
Binary search doesn't fundamentally require a sorted array. It requires a monotonic function over a sorted search space.
1. Search Space: This is the range of all possible answers to the problem. For any valid problem, there's a minimum possible answer and a maximum possible answer. This range of answers is our "search space," and it's inherently sorted.
2. Monotonic Function (The "Check" Function): We need a way to test a potential answer. Let's call this function check(x), where x is a value from our search space. This function must be monotonic. This means it exhibits one of the following properties:
   - If check(x) is true, then for all y > x, check(y) is also true.
   - If check(x) is true, then for all y < x, check(y) is also true.
   This creates a clear boundary in our search space. All values on one side of the boundary make check() return true, and all values on the other side make it return false. The goal of the binary search is to find this boundary, which corresponds to the optimal answer.

The `Sorted Array` statement is essentially true when the search space is the array's indices and the goal is to find a specific value within that array. In this classic scenario:
1. Search Space: The indices of the array, from 0 to n-1. This is an ordered set of integers.
2. Monotonic Function: The comparison with the target value. Let's define a function check(index) = (array[index] < target).
   - If the array is sorted, this check function is guaranteed to be monotonic. It will be true for a block of indices and then false for the rest.
   - [ T, T, T, ..., T, F, F, F ]
   - Without the array being sorted, the results of check(index) would be random ([T, F, T, T, F, ...]), and we couldn't reliably eliminate half the search space.
So for the problem *find value x in collection A*, the sorted property is what creates the required monotonicity and the true requirement is a **monotonic check function over an ordered search space**.
- **Classic BS**: The search space is array indices. The sorted nature of the array *provides* the monotonicity. 
- **BS on the Answer**: The search space is the range of possible answers (e.g., capacity, speed). The inherent logic of the problem ("if speed k works, k+1 must also work") *provides* the monotonicity.


### When can we use Binary Search:
- When need to search for an index or element in a space and if the space is unordered, we can always sort it first before applying Binary Search.
> **Use Case 1: Searching in a Sorted Collection.** This is the classic application. The sorted nature of the data is what allows us to create a monotonic check (e.g., `is array[mid] < target?`) that reliably eliminates half of the search space (the indices).

- If we can discover some kind of **monotonicity** in the problem, we can use binary search. This is the core principle.
> **Use Case 2: Binary Search on the Answer.** If the problem asks for an optimal value (e.g., minimum speed, maximum capacity), we can often binary search over the range of possible answers. We need to define a `check(x)` function that verifies if an answer `x` is possible. This function must be monotonic: if `check(x)` is true, then `check(x+1)` is also true (or vice-versa). This creates an *implicitly sorted* search space of `[False, False, ..., True, True]`.

> **Use Case 3: Other Logical Eliminations.** Some problems don't have a simple monotonic property but still allow you to discard half the search space with each check. A good example is 162. Find Peak Element, where by comparing `nums[mid]` with its neighbor, you can always determine which half is guaranteed to contain a peak.


### How to Spot This Pattern
Ask yourself these three questions when looking at a problem:
1. Is the problem asking for an optimal value (min/max) of some variable x? (e.g., minimum capacity, minimum speed, maximum rows)
2. Can I easily determine the range of possible answers for x? (e.g., [max_weight, sum_weight], [1, max_pile])
3. If I pick a random value mid from this range, can I write a function check(mid) that tells me if mid is a "valid" or "too high/low" answer? Does this function have a monotonic property? (e.g., "If capacity C works, will C+1 also work?")

If you can answer "yes" to all three, you can almost certainly apply "Binary Search on the Answer"! You've uncovered a very powerful and common competitive programming pattern. Great job on your analysis
### Template
```go
package BinarySearch

/* It's the basic Binary_Search template: 704. Binary_Search: https://leetcode.com/problems/binary-search/description/ */
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

### Examples
```go
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

/* Further variation: Sorted array that may contain repeated values */
func searchVariation(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		// "equal" is for handling repeating numbers
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// "left" is the insert position no matter if there are any repeating numbers
	return left
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

/*
1011. Capacity To Ship Packages Within D Days: https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
A conveyor belt has packages that must be shipped from one port to another within days days.

The ith package on the conveyor belt has a weight of weights[i]. Each day, we load the ship with packages on the conveyor belt (in the order given by weights).
We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within days days.
*/
func shipWithinDays(weights []int, days int) int {
	return BinarySearch.ShipWithinDays(weights, days)
}

/*
441. ArrangingCoins: https://leetcode.com/problems/arranging-coins/description/
You have n coins, and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins.
The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.
*/
func arrangeCoins(n int) int {
	return BinarySearch.ArrangeCoinsBS(n)
}
```


