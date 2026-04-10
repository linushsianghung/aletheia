package HashTable

// https://leetcode.com/problems/3sum-with-multiplicity/description/
// Reference: https://chatgpt.com/share/69d5c964-b5f0-8322-80ac-f803598da82e
/*
Given an integer array arr, and an integer target, return the number of tuples i, j, k such that i < j < k and arr[i] + arr[j] + arr[k] == target.

As the answer can be very large, return it modulo 109 + 7.

Analysis:
- Mental Model Instead of "Which indices form valid triplets?", we think "How many ways can values form valid triplets?"
- 3Sum variant
	- Classic 3Sum → avoid duplicates
	- This problem → embrace duplicates and count them
- Introduce frequency counting: How many ways to pick values (i, j, k) such that i + j + k = target?


This problem tests whether you can:
1. Recognize constraints → small value range ⇒ use counting
2. Convert index problem → value problem
3. Handle duplicates correctly → combinatorics (C(n,2), C(n,3))
*/
func threeSumMulti(arr []int, target int) int {
	const mod = 1_000_000_007

	// Step 1: Count frequency: how many times x appears
	count := make([]int, 101)
	for _, v := range arr {
		count[v]++
	}

	result := 0
	// Step 2: Enumerate i, j
	for i := 0; i <= 100; i++ {
		if count[i] == 0 {
			continue
		}
		for j := i; j <= 100; j++ {
			if count[j] == 0 {
				continue
			}

			k := target - i - j
			if k < 0 || k > 100 || k < j {
				continue
			}
			if count[k] == 0 {
				continue
			}

			// Case 1: i < j < k
			if i < j && j < k {
				result += count[i] * count[j] * count[k]
			} else if i == j && j < k {
				// Case 2: i == j < k
				result += count[i] * (count[i] - 1) / 2 * count[k]
			} else if i < j && j == k {
				// Case 3: i < j == k
				result += count[i] * count[j] * (count[j] - 1) / 2
			} else if i == j && j == k {
				// Case 4: i == j == k
				result += count[i] * (count[i] - 1) * (count[i] - 2) / 6
			}

			result %= mod
		}
	}

	return result
}
