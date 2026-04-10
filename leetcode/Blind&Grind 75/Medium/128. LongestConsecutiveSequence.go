package Medium

// LongestConsecutive https://leetcode.com/problems/longest-consecutive-sequence
// Reference:
// - https://leetcode.com/problems/longest-consecutive-sequence/solutions/41057/simple-on-with-explanation-just-walk-eac-ovr2/
// - https://www.youtube.com/watch?v=P6RZZMu_maU
/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.

***Failed in specific case: Time Limit Exceeded***
It might be because the limitation of language. The Java implementation of the same algorithm below can do the job
*/
func LongestConsecutive(nums []int) int {
	note := make(map[int]bool)
	for _, num := range nums {
		note[num] = true
	}

	maxLen := 0
	for _, num := range nums {
		// The head of the sequence is there is no consecutive number before this number. If that's the case, just start to count the length
		if !note[num-1] {
			length := 1

			for note[num+length] {
				length++
			}
			maxLen = max(maxLen, length)
		}
	}

	return maxLen
}

/*
class Solution {
    public int longestConsecutive(int[] nums) {
        Set<Integer> set = new HashSet<>();
        for(int n : nums) {
            set.add(n);
        }

        int maxLen = 0;
        for(int n : set) {
            if(!set.contains(n - 1)) {
                int len = 1;

                while(set.contains(n+len)) {
                    len++;
                }

                maxLen = Math.max(maxLen, len);
            }
        }
        return maxLen;
    }
}
*/
