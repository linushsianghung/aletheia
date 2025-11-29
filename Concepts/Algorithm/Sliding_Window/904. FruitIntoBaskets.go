package Sliding_Window

// https://leetcode.com/problems/fruit-into-baskets/description/
/*
You are visiting a farm that has a single row of fruit trees arranged from left to right.
The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.
You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

- You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
- Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
- Once you reach a tree with fruit that cannot fit in your baskets, you must stop.

Given the integer array fruits, return the maximum number of fruits you can pick.
*/
func totalFruit(fruits []int) int {
	winStart, maxCount := 0, 0
	// Window State: types of fruit with its amount currently we have
	note := make(map[int]int)

	for winEnd, fruit := range fruits {
		note[fruit]++

		// Dynamic-Size Sliding Window: based on the types of fruit we have
		for len(note) > 2 {
			note[fruits[winStart]]--
			if note[fruits[winStart]] == 0 {
				delete(note, fruits[winStart])
			}
			winStart++
		}

		maxCount = max(maxCount, winEnd-winStart+1)
	}

	return maxCount
}

func totalFruitExercise(fruits []int) int {
	return 0
}
