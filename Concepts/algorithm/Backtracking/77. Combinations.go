package Backtracking

// https://leetcode.com/problems/combinations/description/
// Ref:
// - https://leetcode.com/problems/combinations/solutions/844096/backtracking-cheatsheet-simple-solution/
// - https://leetcode.com/problems/combinations/solutions/1141903/clear-and-simple-explanation-with-intuition-100-faster/
/*
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.

Discussion:
DFS:
- https://leetcode.com/problems/combinations/solutions/27006/a-template-to-those-combination-problems/
- https://leetcode.com/problems/combinations/solutions/429526/general-backtracking-questions-solutions-in-python-for-reference/
*/
func combine(n int, k int) [][]int {
	/*
		Given an empty array, the task is to add numbers between 1 to n to the array up to the size of k.
		We could model the step to add a number as a recursion function (i.e. backtrack() function).
	*/
	result := make([][]int, 0)
	backtrackCombinations(&result, make([]int, 0), 1, n, k)
	return result
}

func backtrackCombinations(result *[][]int, processor []int, start, n, k int) {
	/* The backtracking would be triggered at the points where the decision space is completed i.e., start is 9 or when the size of becomes k. */
	// Base Case => Every problem of backtracking has some base case which tells us at which point we have to stop with the recursion process.
	if len(processor) == k {
		// Create a new copy of processor slice to prevent from using the same underlying array,
		// so that the result won't be modified after removing and appending a new element from/to process slice
		p := make([]int, k)
		copy(p, processor)
		*result = append(*result, p)
		return
	}

	/*
		Technically, we have 9 candidates at hand to add to the array. Yet we want to consider solutions that lead to a valid case (i.e. is_valid(candidate)).
		Here the validity is determined by whether the number is repeated or not. Since in the loop, we have iterated from start to n+1,
		the numbers before start are already visited and cannot be added to the slice. Hence, we don't arrive at an invalid case.
	*/
	// Iterate through all possible candidates
	for i := start; i <= n; i++ {
		/* Then, among all the suitable candidates, we add different numbers using append(i). Later we can revert our decision by re-slicing, so that we could try out the other candidates. */
		processor = append(processor, i)
		// Or create new copy before passing processor slice to next level of backtrack function
		// p := make([]int, len(processor))
		// copy(p, processor)
		backtrackCombinations(result, processor, i+1, n, k)
		processor = processor[:len(processor)-1]
	}
}
