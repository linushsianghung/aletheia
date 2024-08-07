package String

// https://leetcode.com/problems/excel-sheet-column-number
/*
Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
*/
func titleToNumber(columnTitle string) int {
	if len(columnTitle) == 0 {
		return 0
	}

	sum := 0
	for _, s := range columnTitle {
		sum *= 26
		sum += int(s) - 'A' + 1
	}

	return sum
}
