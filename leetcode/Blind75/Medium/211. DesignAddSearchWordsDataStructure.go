package Medium

// https://leetcode.com/problems/design-add-and-search-words-data-structure
/*
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

WordDictionary() Initializes the object.
void addWord(word) Adds word to the data structure, it can be matched later.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
*/

type WordDictionary struct {
	data []string
}

func Constructor() WordDictionary {
	return WordDictionary{
		data: make([]string, 0),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.data = append(this.data, word)
}

func (this *WordDictionary) Search(word string) bool {
dataloop:
	for _, d := range this.data {
		if len(d) != len(word) {
			continue
		}

		for i, w := range word {
			if w == '.' || w == rune(d[i]) {
				continue
			} else {
				continue dataloop
			}
		}

		return true
	}

	return false
}
