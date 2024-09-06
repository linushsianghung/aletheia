package data_structure

// MyHashSet https://leetcode.com/problems/design-hashset/description/
/*
Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:

void add(key) Inserts the value key into the HashSet.
bool contains(key) Returns whether the value key exists in the HashSet or not.
void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.
*/
type MyHashSet struct {
	innerSlice []int
}

func HashSetConstructor() MyHashSet {
	tempSlice := make([]int, 1000001)
	for i := range tempSlice {
		tempSlice[i] = -1
	}
	return MyHashSet{
		innerSlice: tempSlice,
	}
}

func (this *MyHashSet) Add(key int) {
	this.innerSlice[key] = key
}

func (this *MyHashSet) Remove(key int) {
	this.innerSlice[key] = -1
}

func (this *MyHashSet) Contains(key int) bool {
	return this.innerSlice[key] != -1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
