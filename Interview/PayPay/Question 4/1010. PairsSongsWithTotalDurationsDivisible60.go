package Question_4

// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/description/
// Reference: https://chatgpt.com/share/69d4b4f8-ed30-8324-81cf-1d50539f22cf
/*
You are given a list of songs where the ith song has a duration of time[i] seconds.

Return the number of pairs of songs for which their total duration in seconds is divisible by 60. Formally, we want the number of indices i, j such that i < j with (time[i] + time[j]) % 60 == 0.

Analysis:
For any two songs a and b: (a + b) % 60 == 0
That means: a % 60 + b % 60 == 60 OR == 0

So if:
- r = time[i] % 60
- Then we need a previous song with remainder: complement = (60 - r) % 60
*/
func numPairsDivisibleBy60(time []int) int {
	remainder := make([]int, 60) // remainder frequency
	count := 0

	for _, t := range time {
		r := t % 60
		count += remainder[(60-r)%60]
		remainder[r]++
	}

	return count
}
