package testdome

import (
	"fmt"
	"strings"
)

func isAnagram(a string, b string) bool {
	var alphabet [26]int

	strings.ToLower(a)
	strings.ToLower(b)
	for _, c := range a {
		position := int(c - 'a')
		alphabet[position]++
	}

	for _, c := range b {
		position := int(c - 'a')
		alphabet[position]--
	}

	for _, n := range alphabet {
		if n != 0 {
			return false
		}
	}

	return true
}

func Anagram() {
	a := "cde"
	b := "abc"
	fmt.Println(isAnagram(a, b))
}
