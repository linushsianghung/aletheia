package Backtracking

// https://leetcode.com/problems/palindrome-partitioning/description/
// Reference: https://www.youtube.com/watch?v=sz1qaKt0KGQ
/*
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome all partitioning of s.

Analysis:
## Template
All backtracking problems are composed by these three steps: `choose`, `explore`, `un-choose`.
So for each problem, you need to know:

- Choose what? For this problem, we choose each substring.
- How to explore? For this problem, we do the same thing to the remained string.
- Un-Choose; Do the opposite operation of choosing.

Let's take this problem as an example:
1. Define *backtrack()*: Usually we need a helper function in a backtracking problem to accept more parameters.
2. Parameters: Usually we need the following parameters
>	The object that you are working on: for this problem is String s.
>	A start index or an end index which indicates which part you are working on: For this problem, we use substring to indicate the start index.
>	A processor, to remember current choose and then do un-choose: For this problem, we use List<String> processor.
>	A final result, to remember the final result. Usually when we add, we use 'result.add(new ArrayList<>(step))' instead of 'result.add(step)',
>	since the step is passed by reference. We will modify step later, so we need to copy it and add the copy to the result;
3. Base case: The base case defines when to add processor into the result, and when to return.
4. Use for-loop: Usually we need a for loop to iterate though the input String s, so that we can choose all the options.
5. Choose: In this problem, if the substring of s is palindrome, we add it into the step, which means we choose this substring.
6. Explore: In this problem, we want to do the same thing to the remaining substring. So we recursively call our function.
7. Un-Choose: We draw back, remove the chosen substring in order to try other options.
*/
func partition(s string) [][]string {
	return nil
}
