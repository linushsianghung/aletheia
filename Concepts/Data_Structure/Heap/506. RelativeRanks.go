package Heap

import (
	"container/heap"
	"strconv"
)

// https://leetcode.com/problems/relative-ranks/
/*
You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition. All the scores are guaranteed to be unique.

The athletes are placed based on their scores, where the 1st place athlete has the highest score, the 2nd place athlete has the 2nd highest score, and so on. The placement of each athlete determines their rank:

- The 1st place athlete's rank is "Gold Medal".
- The 2nd place athlete's rank is "Silver Medal".
- The 3rd place athlete's rank is "Bronze Medal".
- For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athlete's rank is "x").

Return an array answer of size n where answer[i] is the rank of the ith athlete.
*/
func findRelativeRanks(score []int) []string {
	note := make(map[int]int)
	for i, s := range score {
		note[s] = i
	}

	scoreHeap := &IntHeap{}
	heap.Init(scoreHeap)

	for score := range note {
		heap.Push(scoreHeap, score)
	}

	answer := make([]string, len(score))
	rank := 1

	for scoreHeap.Len() > 0 {
		score := heap.Pop(scoreHeap).(int)
		index := note[score]

		switch rank {
		case 1:
			answer[index] = "Gold Medal"
		case 2:
			answer[index] = "Silver Medal"
		case 3:
			answer[index] = "Bronze Medal"
		default:
			answer[index] = strconv.Itoa(rank)
		}
		rank++
	}

	return answer
}
