package data_structure

// MyHashMap https://leetcode.com/problems/design-hashmap/description/
/*
Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
*/
type MyHashMap struct {
	innerSlice []int
}

func HashMapConstructor() MyHashMap {
	tempSlice := make([]int, 1000001)
	for i := range tempSlice {
		tempSlice[i] = -1
	}
	return MyHashMap{
		innerSlice: tempSlice,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.innerSlice[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.innerSlice[key]
}

func (this *MyHashMap) Remove(key int) {
	this.innerSlice[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
