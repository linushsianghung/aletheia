package Array

import "sort"

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description
// Ref: https://leetcode.com/problems/intersection-of-two-arrays-ii/solutions/282372/java-solution-with-all-3-follow-up-questions/
/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear
as many times as it shows in both arrays, and you may return the result in any order.
*/
func intersect(nums1 []int, nums2 []int) []int {
	// return intersectMap(nums1, nums2)
	return intersect2Pointers(nums1, nums2)
}

/*
Algorithm:
Using a map to save the number of each element in nums1 array then going through each element in nums2 array.
If the element can be found in the map, add that element in the result array and update the map.

Complexity Analysis:
- Time Complexity: O(N + M), O(N) for iterate one of the array to create a hashmap and O(M) to iterate the other array.
- Space Complexity: O(N) to store such hashmap.
*/
func intersectMap(nums1 []int, nums2 []int) []int {
	note, result := make(map[int]int), make([]int, 0)

	for _, num := range nums1 {
		note[num]++
	}

	for _, num := range nums2 {
		if v := note[num]; v > 0 {
			result = append(result, num)
			note[num]--
		}
	}

	return result
}

func intersectMapExercise(nums1 []int, nums2 []int) []int {
	return nil
}

// Follow up:
// What if the given array is already sorted? How would you optimise your Algorithm?
/*
Complexity Analysis:
- Time Complexity: O(max(N, M)). Worst case, for example, would be nums1 = {100}, and nums2 = {1, 2, ..., 100 }. We will always iterate the longest array
*/
func intersect2Pointers(nums1 []int, nums2 []int) []int {
	// Pre-Step: O(MlogM + NlogN)
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		num1, num2 := nums1[i], nums2[j]

		if num1 < num2 {
			i++
		} else if num1 > num2 {
			j++
		} else {
			result = append(result, num1)
			i++
			j++
		}
	}

	return result
}

// What if nums1's size is small compared to nums2's size? Which Algorithm is better?
/*
Complexity Analysis:
- Time Complexity: O(K(logN) + N). Plus N is worst case scenario which you have to linear scan every element in A. But on average, that shouldn't be the case. so I'd say the Time complexity is O(K(logN) + c)
*/

// What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

/*
Analysis:
1. Hash Map Approach:
   - Concept: We count frequencies of each number in one array (preferably the smaller one) and then iterate through the second array to find commonalities.
   - Time Complexity: O(N + M) where N and M are the lengths of the two arrays. We traverse each array once.
   - Space Complexity: O(min(N, M)) because we store the elements of the smaller array in the hash map to minimize memory usage.

2. Two Pointers Approach:
   - Concept: If the arrays are sorted, we can use two pointers to find intersections in a single linear scan without extra hash table overhead.
   - Time Complexity: O(N log N + M log M) if sorting is required. If already sorted, it is O(N + M).
   - Space Complexity: O(log N + log M) to O(N + M) depending on the sorting implementation (e.g., Dual-Pivot Quicksort in Java).

3. Interview Suggestions (Follow-ups):
   - Small vs Large: If `nums1` is much smaller than `nums2`, use the Map approach on `nums1`. It keeps the space complexity low.
   - Disk/Limited Memory: If `nums2` is on disk and doesn't fit in memory, but `nums1` does, load `nums1` into a Map. If neither fits, use an external sort on both and then use the two-pointer approach to stream elements.

Java Implementation:

import java.util.*;

class Solution {
    // Method 1: Hash Map
    public int[] intersectMap(int[] nums1, int[] nums2) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums1) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        List<Integer> intersection = new ArrayList<>();
        for (int num : nums2) {
            int count = map.getOrDefault(num, 0);
            if (count > 0) {
                intersection.add(num);
                map.put(num, count - 1);
            }
        }

        // Convert List to primitive array
        int[] result = new int[intersection.size()];
        for (int i = 0; i < intersection.size(); i++) {
            result[i] = intersection.get(i);
        }
        return result;
    }

    // Method 2: Two Pointers (Best if sorted)
    public int[] intersect2Pointers(int[] nums1, int[] nums2) {
        Arrays.sort(nums1);
        Arrays.sort(nums2);

        int i = 0, j = 0, k = 0;
        // We can reuse nums1 to store the result to save extra space
        while (i < nums1.length && j < nums2.length) {
            if (nums1[i] < nums2[j]) {
                i++;
            } else if (nums1[i] > nums2[j]) {
                j++;
            } else {
                nums1[k++] = nums1[i++];
                j++;
            }
        }
        return Arrays.copyOfRange(nums1, 0, k);
    }

    // Driver method matching LeetCode signature
    public int[] intersect(int[] nums1, int[] nums2) {
        // Choose one based on constraints discussed
        return intersectMap(nums1, nums2);
    }
}
*/
