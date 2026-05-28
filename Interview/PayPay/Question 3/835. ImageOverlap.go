package Question_3

// https://leetcode.com/problems/image-overlap/
// Reference:
// - https://www.youtube.com/watch?v=fkSq0QE5C_0
// - https://chatgpt.com/share/69d1d64e-6f3c-8320-a40b-c6976368b0a9
/*
You are given two images, img1 and img2, represented as binary, square matrices of size n x n. A binary matrix has only 0s and 1s as values.

We translate one image however we choose by sliding all the 1 bits left, right, up, and/or down any number of units. We then place it on top of the other image. We can then calculate the overlap by counting the number of positions that have a 1 in both images.

Note also that a translation does not include any kind of rotation. Any 1 bits that are translated outside of the matrix borders are erased.

Return the largest possible overlap.
*/
/* Complexity Analysis
   Time Complexity: O(N^2 + L1 * L2)
   - N^2 to scan the matrices and find coordinates of '1's.
   - L1 * L2 to calculate all possible shift vectors, where L1 and L2 are the number of 1s in img1 and img2.
   - In the worst case (all 1s), L1 = L2 = N^2, making it O(N^4). However, for sparse matrices, this is much faster.
   - Given N=30, N^4 is 810,000, which is well within limits for Go.

   Space Complexity: O(L1 + L2 + (L1 * L2))
   - O(L1 + L2) to store the coordinates.
   - O(L1 * L2) in the worst case for the hash map to store unique shift vectors.
*/
func largestOverlap(img1 [][]int, img2 [][]int) int {
	size := len(img1)

	list1, list2 := make([][2]int, 0), make([][2]int, 0)
	for i := range size {
		for j := range size {
			if img1[i][j] == 1 {
				list1 = append(list1, [2]int{i, j})
			}
			if img2[i][j] == 1 {
				list2 = append(list2, [2]int{i, j})
			}
		}
	}

	// Since N is small (up to 30), we can represent the shift (dx, dy) as a single integer to potentially speed up map lookups.
	// A simple transformation: dx * 100 + dy (since range is roughly -30 to 30)
	note := make(map[int]int)
	//note := make(map[[2]int]int)
	for _, anchor := range list1 {
		for _, target := range list2 {
			x := target[0] - anchor[0]
			y := target[1] - anchor[1]
			
			shiftKey := x*100 + y
			note[shiftKey]++
			//note[[2]int{x, y}]++
		}
	}

	maxCount := 0
	for _, val := range note {
		maxCount = max(maxCount, val)
	}

	return maxCount
}
