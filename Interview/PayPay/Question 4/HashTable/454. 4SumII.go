package HashTable

// https://leetcode.com/problems/4sum-ii/description/
// Reference: https://leetcode.com/problems/4sum-ii/solutions/1740606/going-from-on4-on3-on2-javac-by-hi-malik-2kxh/
/*
Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	return 0
}

func fourSumCount2HashMap(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	note := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			note[num1+num2]++
		}
	}

	count := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			sum := num3 + num4

			if val, ok := note[-sum]; ok {
				count += val
			}
		}
	}

	return count
}

func fourSumHashMap(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	note := make(map[int]int)
	for _, num := range nums4 {
		note[num]++
	}

	count := 0
	for _, i := range nums1 {
		for _, j := range nums2 {
			for _, k := range nums3 {
				sum := i + j + k
				if val, ok := note[-sum]; ok {
					count += val
				}
			}
		}
	}

	return count
}

func fourSumCountBruteForce(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0

	for _, i := range nums1 {
		for _, j := range nums2 {
			for _, k := range nums3 {
				for _, l := range nums4 {
					if i+j+k+l == 0 {
						count++
					}
				}
			}
		}
	}

	return count
}
