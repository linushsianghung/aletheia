package Backtracking

/*
# Backtracking
Ref:
- The Backtracking Blueprint: https://www.youtube.com/watch?v=Zq4upTEaQyM
- Java Template: https://leetcode.com/problems/subsets/solutions/27281/a-general-approach-to-backtracking-questions-in-java-subsets-permutations-combination-sum-palindrome-partitioning/
- NeetCode: https://www.youtube.com/playlist?list=PLot-Xpze53lf5C3HSjCnyFghlW0G1HHXo
- How to calculate runtime of backtracking algorithm: https://leetcode.com/discuss/interview-question/3055778/how-to-calculate-runtime-of-backtracking-algorithm-for-interview

Discussion:
- Leetcode Pattern 3 | Backtracking: https://medium.com/leetcode-patterns/leetcode-pattern-3-backtracking-5d9e5a03dc26

## Definition:
- Backtracking is a general algorithm for finding solutions to some computational problems, notably constraint satisfaction problems, that incrementally builds candidates
to the solutions, and abandons a candidate ("backtracks") as soon as it determines that the candidate cannot possibly be completed to a valid solution.
- Backtracking can be seen as an optimised way to brute force. Brute force approaches evaluate every possibility. In backtracking, you stop evaluating a possibility
as soon as it breaks some constraint provided in the problem, take a step back and keep trying other possible cases, see if those lead to a valid solution.
- Due to this backtracking behaviour, the backtracking algorithms are often much faster than the brute-force search algorithm since it eliminates many unnecessary explorations.

## Backtracking Problems:
> You are explicitly asked to return a collection of all answers.
> You are concerned with what the actual solutions are rather than say the most optimum value of some parameter. (if it were the latter, it’s most likely DP or greedy).

## Discussion:
- Difference between back tracking and dynamic programming: https://stackoverflow.com/questions/3592943/difference-between-back-tracking-and-dynamic-programming
- How do I decide between Dynamic Programming vs. Backtracking: https://www.reddit.com/r/leetcode/comments/ntuycc/how_do_i_decide_between_dynamic_programming_vs/
- Backtracking v.s. DFS:
	- DFS + not visiting an invalid node => Backtracking.
	- DFS + not visiting node twice => Dynamic Programming. [let's ignore tabular for now]
*/

type container struct{} // Something can hold values, like []int
type element struct{}

func (c container) add(e element)               {}
func (c container) remove(e element)            {}
func (c container) merge(e container) container { return e }

func backtrackingTemplate(source []element) [][]element {
	return backtrack(source)

	// factors
	// var x, y int
	// var result [][]element
	// backtrackHelper(&result, []element{}, source, 0, x, y)
	// return result
}

/*
* Choice - Think about the choice you are making at every step
* Constraints - What rules the choice must follow
* Goal - What base case needs to be reached?
 */
/* Approach 1.
By using a local function with closure, the result variable can be accessed directly instead of passing as a pointer, which might not be easy
to be understood and maintained
*/
func backtrack(source []element) [][]element {
	var constraintIsMatch, goalIsMatch bool
	var decisionSpace []element // Based on the source
	result := make([][]element, 0)

	var localRecursiveFunc func(source, processor []element, progress int, factor ...interface{})
	localRecursiveFunc = func(source, processor []element, progress int, factor ...interface{}) {
		// Goal, i.e. Base Case
		if goalIsMatch {
			// Add up the current processor to the result or not, then backtrack
			result = append(result, processor)
			return
		}

		// Choice: all the paths of the Decision Space
		for _, e := range decisionSpace {
			// Constraints: only recursing on a valid path
			if constraintIsMatch {
				processor = append(processor, e) // Considering the current candidate.
				/* Create a new copy of processor slice here to prevent from using the same underlying array,
				so that the result won't be modified after removing and appending a new element from/to process slice */
				p := make([]element, len(processor))
				copy(p, processor)
				localRecursiveFunc(source, p, 0, factor) // Considering & Exploring based on the current candidate further
				processor = processor[:len(processor)-1] // Remove the current candidate => Backtrack
			}
		}
	}

	localRecursiveFunc(source, []element{}, 0)
	return result
}

/*
Approach 2.
Passing the result variable as a pointer is a bit more straightforward, but the code might not be concise as Approach 1
*/
func backtrackHelper(result *[][]element, source, processor []element, progress int, factor ...interface{}) {
	var eIsValid, constraintIsMatch bool
	var decisionSpace []element
	if constraintIsMatch {
		*result = append(*result, processor)
		return
	}

	// Choice: all the paths of the Decision Space
	for _, e := range decisionSpace {
		// Constraints: only recursing on a valid path
		if eIsValid {
			processor = append(processor, e)
			p := make([]element, len(processor))
			copy(p, processor)
			backtrackHelper(result, source, p, 0, factor)
			processor = processor[:len(processor)-1]
		}
	}
}
