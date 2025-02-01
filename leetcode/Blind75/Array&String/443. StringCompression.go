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
	var count int

	for i := 0; i < len(chars); {
		anchor := chars[i]
		repeat := 0

		for i < len(chars) && chars[i] == anchor {
			repeat++
			i++
		}

		chars[count] = anchor
		count++
		if repeat > 1 {
			// The repeat might over 10 times, so use for loop to convert number to characters
			for _, c := range strconv.Itoa(repeat) {
				chars[count] = byte(c)
				count++
			}
		}
	}

	return count
}
