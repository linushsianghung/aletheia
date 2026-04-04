package Question_3

// https://leetcode.com/problems/image-overlap/
// Reference: https://www.youtube.com/watch?v=fkSq0QE5C_0
/*
You are given two images, img1 and img2, represented as binary, square matrices of size n x n. A binary matrix has only 0s and 1s as values.

We translate one image however we choose by sliding all the 1 bits left, right, up, and/or down any number of units. We then place it on top of the other image. We can then calculate the overlap by counting the number of positions that have a 1 in both images.

Note also that a translation does not include any kind of rotation. Any 1 bits that are translated outside of the matrix borders are erased.

Return the largest possible overlap.

Analysis:
https://chatgpt.com/share/69d0bf36-0a0c-8322-86be-2c96d1bb2f0d


*/
func largestOverlap(img1 [][]int, img2 [][]int) int {
	list1, list2 := make([][]int, 0), make([][]int, 0)

	for i := 0; i < len(img1); i++ {
		for j := 0; j < len(img1[0]); j++ {
			if img1[i][j] == 1 {
				list1 = append(list1, []int{i, j})
			}
			if img2[i][j] == 1 {
				list2 = append(list2, []int{i, j})
			}
		}
	}

	note := make(map[[2]int]int)
	for _, anchor := range list1 {
		for _, target := range list2 {
			x := target[0] - anchor[0]
			y := target[1] - anchor[1]

			note[[2]int{x, y}]++
		}
	}

	maxCount := 0
	for _, value := range note {
		maxCount = max(maxCount, value)
	}

	return maxCount
}
