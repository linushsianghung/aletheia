package Array_String

import "strconv"

// https://leetcode.com/problems/string-compression/description/?envId=leetcode-75
/*
Given an array of characters chars, compress it using the following algorithm:
Begin with an empty string s. For each group of consecutive repeating characters in chars:

- If the group's length is 1, append the character to s.
- Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars.
Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.
You must write an algorithm that uses only constant extra space.
*/
func compress(chars []byte) int {
	var anchor int

	for i := 0; i < len(chars); {
		current := chars[i]
		repeat := 0

		for i < len(chars) && chars[i] == current {
			repeat++
			i++
		}

		chars[anchor] = current
		anchor++
		if repeat > 1 {
			// The repeat might over 10 times, so use for loop to convert number to characters
			for _, c := range strconv.Itoa(repeat) {
				chars[anchor] = byte(c)
				anchor++
			}
		}
	}

	return anchor
}

func compressExercise(chars []byte) int {
	return 0
}

// This intuitive implementation is just going through each element 1 by 1, but it has to handle last element specially which results in duplicate code
func compressIntuitively(chars []byte) int {
	anchor, count := 0, 1

	for i := 1; i < len(chars); i++ {
		if chars[i-1] == chars[i] {
			count++
		} else {
			chars[anchor] = chars[i-1]
			anchor++
			if count > 1 {
				for _, c := range strconv.Itoa(count) {
					chars[anchor] = byte(c)
					anchor++
				}
			}

			count = 1
		}
	}

	chars[anchor] = chars[len(chars)-1]
	anchor++
	if count > 1 {
		for _, c := range strconv.Itoa(count) {
			chars[anchor] = byte(c)
			anchor++
		}
	}

	return anchor
}
