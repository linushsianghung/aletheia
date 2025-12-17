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
It's the basic Binary Search Template: 704. Binary_Search: https://leetcode.com/problems/binary-search/description/
```go
package BinarySearch
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
