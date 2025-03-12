package Array_String

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/?envId=leetcode-75
/*
There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has,
and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies,
they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.
*/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, candy := range candies {
		maxCandies = max(maxCandies, candy)
	}

	result := make([]bool, len(candies))
	for i, candy := range candies {
		if candy+extraCandies >= maxCandies {
			result[i] = true
		}
	}

	return result
}

func kidsWithCandiesExercise(candies []int, extraCandies int) []bool {
	return nil
}
