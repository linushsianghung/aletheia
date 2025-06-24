package Trie

// https://leetcode.com/problems/implement-trie-prefix-tree/description/?envId=leetcode-75
/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/

type Trie struct {
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {

}

func (this *Trie) Search(word string) bool {
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}
