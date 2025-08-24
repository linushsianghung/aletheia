package Medium

import "math"

// MaxSubArray https://leetcode.com/problems/maximum-subarray/
// Ref:
// - Back To Back SWE: [Max Contiguous Subarray Sum - Cubic Time To Kadane's Algorithm](https://www.youtube.com/watch?v=2MmGzdiKR9Y)
// - NeetCode: [Maximum Subarray](https://www.youtube.com/watch?v=5WZl3MMT0Eg)
// - https://leetcode.com/problems/maximum-subarray/solutions/1595195/c-python-7-simple-solutions-w-explanation-brute-force-dp-kadane-divide-conquer/
/*
Given an integer array nums, find the sub-array with the largest sum, and return its sum.

Analysis:
Key Word => Contiguous

1. O(n^3) time complexity: computing all possibilities of sub-array and starting from the head everytime
2. O(n^2) time complexity: using Sliding_Window technique when computing all the sub-array
3. O(n^1) time complexity: using Kadane's Algorithm (Dynamic_Programming):
	- Inverse the computation to define the sub-problem: what is the maximum sum of contiguous sub-array that end at the index i (ms(i))
	- At each end index of the iteration, what are the choices we have for each element to produce the maximum sum
		- choice 1: ms(i) is index i itself
		- choice 2: ms(i) is the sum of ms(i-1) + i
	- Then max(choice 1, choice 2)


- O(N^3):
Brute-Force for calculate every single sub-array to find out the maximum sum
	  2 -1 -3 4 1 -2 5 1 -3 2
sum:  2~> 1
	  2~~> -2
	  2~~~> 2
	  .........
		   -3~> 7
		   -3~~> 8
		   -3~~~> 6
			..........
- O(N^2):
Just add new element to the currentBestSum to eliminate duplicated calculation. It also provides another perspective of the problem, that is,
at any element i, the maximum sum should be the sum of one of the previous sub-array ending at element i-1 plus current element (like below image),
which is supposed to be the max one
	  2 -1 -3 4 1 -2 5 1 -3 2
sum:	  -1<~~
         0<~~~~
	   2<~~~~~~
	 ..........

- O(N):
How can we know which previous sub-array ending at element i-1 is the max one? Basically it's not that easy to figure it out. But if we can find out
the best sum at element i-1, definitely we can get better choice of current element i, namely currentBestSum = max(nums[i], currentBestSum + nums[i])
			 			 2 -1 -3 4 1 -2 5 1 -3 2
evaluation:  max(-1, 1)  ~~> 1
	  		 max(-3, -2) ~~~> -2
		  	 max(4, 2)   ~~~~> 4
			 ..........
*/
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum, currentBestSum := math.MinInt32, math.MinInt32
	for _, num := range nums {
		/* Kadane's Algorithm (Dynamic_Programming):
		   For each element [i], the value of element [i-1] is the best solution of subarray ending at that point.
		   So just consider using the value of element [i] itself or includes the previous best value from element [i-1]

		   currentBestSum = max(nums[i], currentBestSum + nums[i])
		*/
		currentBestSum = max(num, currentBestSum+num)
		maxSum = max(maxSum, currentBestSum)
	}

	return maxSum
}

func maxSubArrayExercise(nums []int) int {
	return 0
}
