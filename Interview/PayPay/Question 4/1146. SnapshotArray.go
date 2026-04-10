package Question_4

// https://leetcode.com/problems/snapshot-array/description/
/*
Implement a SnapshotArray that supports the following interface:

SnapshotArray(int length) initializes an array-like data structure with the given length. Initially, each element equals 0.
void set(index, val) sets the element at the given index to be equal to val.
int snap() takes a snapshot of the array and returns the snap_id: the total number of times we called snap() minus 1.
int get(index, snap_id) returns the value at the given index, at the time we took the snapshot with the given snap_id
*/

type SnapshotArray struct {
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{}
}

func (this *SnapshotArray) Set(index int, val int) {

}

func (this *SnapshotArray) Snap() int {
	return 0
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
