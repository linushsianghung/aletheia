package Array_String

// https://leetcode.com/problems/can-place-flowers/description/?envId=leetcode-75
// Reference: https://leetcode.com/problems/can-place-flowers/solutions/103883/java-very-easy-solution/
/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

Analysis:
If there are zeroCount zeroes in between two 1s, then how many 1s can we place in those zeroes without violating the given condition?
The Answer is (zeroCount-1)/2.

The only cases this doesn't apply are when there are zeroes(1 or more)
- At the beginning of the array.
- At the end of the array.
For these 2 cases, the number of 1s that we can place is zeroCount/2. But to generalize the algorithm and to simplify code inside loop,
zeroCount has been initialized to 1 for the first time and result += (zeroCount-1)/2 effectively becomes result += zeroCount/2 for the case 1.
For case 2, result is updated outside the loop, again by zeroCount/2 times.

Finally, we check if the number of possible 1s that we can place is greater than or equal to n. If so, we return true else false.
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	return canPlaceFlowersSmart(flowerbed, n)
}

func canPlaceFlowersSmart(flowerbed []int, n int) bool {
	zeroCount, flowers := 1, 0

	for _, pot := range flowerbed {
		if pot == 0 {
			zeroCount++
		} else {
			flowers += (zeroCount - 1) / 2
			zeroCount = 0
		}
	}

	flowers += zeroCount / 2
	return flowers >= n
}

func canPlaceFlowersSmartExercise(flowerbed []int, n int) bool {
	return false
}

func canPlaceFlowersAlt(flowerbed []int, n int) bool {
	zeroCount := 0

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
				zeroCount++
			}
		}
	}

	return zeroCount >= n
}
