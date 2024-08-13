package Array_String

// https://leetcode.com/problems/can-place-flowers/description/?envId=leetcode-75
/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			prevEmpty, nextEmpty := false, false

			if i == 0 || flowerbed[i-1] == 0 {
				prevEmpty = true
			}
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
				nextEmpty = true
			}

			if prevEmpty && nextEmpty {
				flowerbed[i] = 1
				count++
			}
		}
	}

	return count >= n
}
