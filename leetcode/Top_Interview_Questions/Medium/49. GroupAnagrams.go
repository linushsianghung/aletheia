package Medium

import "sort"

// GroupAnagrams https://leetcode.com/problems/group-anagrams/
/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Analysis:
Time Complexity: O(N * KlogK)
1. Outer Loop (N): The loop for _, str := range strs runs N times.
2. Inner Operation (Sorting): Inside the loop, you sort a string of length K.
	- Sorting takes O(KlogK)
	- Map access and hashing the string takes $O(K)$.
	- The dominant term inside the loop is the sorting: O(KlogK).
3. Total Complexity: $$N \times O(KlogK) = O(N * KlogK)
If you assume the string length K is small and constant (e.g., always less than 100), you could argue it approaches O(N), but strictly speaking, it depends on both variables.
*/
// GroupAnagrams implements the O(N * K) approach using character counts.
// In Go, arrays (like [26]int) can be used as map keys so instead of sorting (KlogK), we count the frequency of each character (which costs K).
// "eat" -> [1, 0, 0, 0, 1, ... 1 ...] (1 'a', 1 'e', 1 't')
// "tea" -> [1, 0, 0, 0, 1, ... 1 ...] (Same array)
func GroupAnagrams(strs []string) [][]string {
	// Key is an array of 26 integers (for a-z). Arrays are comparable in Go and can be map keys.
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, char := range str {
			count[char-'a']++
		}
		groups[count] = append(groups[count], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func groupAnagramsExercise(strs []string) [][]string {
	return nil
}

func GroupAnagramsNaive(strs []string) [][]string {
	note := make(map[string][]string)

	for _, str := range strs {
		sortedStr := []rune(str)
		sort.Slice(sortedStr, func(i, j int) bool {
			return sortedStr[i] < sortedStr[j]
		})

		key := string(sortedStr)
		note[key] = append(note[key], str)
		//if anagram, ok := note[string(sortedStr)]; ok {
		//	note[string(sortedStr)] = append(anagram, str)
		//} else {
		//	note[string(sortedStr)] = []string{str}
		//}
	}

	result := make([][]string, 0, len(note))
	for _, anagram := range note {
		result = append(result, anagram)
	}

	return result
}

/*
Java Implementation:

Analysis:
1. Concept Introduction:
   The core idea of grouping anagrams is finding a "Canonical Form" or a "Signature" for each string.
   If two strings are anagrams, they must produce the same signature.
   - Analogy: Imagine you have a pile of various LEGO sets mixed together. To group them, you could either
     build the model (Sorting) or count how many of each specific brick type you have (Frequency Counting).
     If two piles have the exact same count of blue, red, and yellow bricks, they belong to the same set.

2. How to reach the Optimized Solution during an Interview:
   - Step 1 (Brute Force): Compare every string with every other string ($O(N^2 * K)$). This is inefficient.
   - Step 2 (Sorting): Realize that sorting an anagram results in the same string. Using a Map<String, List>
     where the key is the sorted string gives $O(N * K \log K)$.
   - Step 3 (Frequency Counting): Ask yourself: "Can I represent the character distribution faster than sorting?"
     Since we only have lowercase English letters, a fixed-size array of 26 integers can represent the count.
     This brings the per-string cost down from $O(K \log K)$ to $O(K)$.
   - Step 4 (Java Specific Constraint): In Java, `int[]` is not a valid key for a `HashMap` because it doesn't
     override `equals()` and `hashCode()` for content. We solve this by converting the array into a
     delimited String (e.g., "#1#0#2...") to use as the map key.

3. Complexity Analysis:
   - Time Complexity: O(N * K)
     - We iterate through N strings.
     - For each string of length K, we iterate through its characters to count them: O(K).
     - We build a string key from the 26-count array: O(1) as the array size is constant.
     - Overall: O(N * K).
   - Space Complexity: O(N * K)
     - We store all strings in the HashMap. In the worst case (no anagrams), the map will store N keys,
       each being a representation of the character counts.

4. Interview Suggestion:
   - Start by mentioning the Sorting approach as it's intuitive.
   - Propose the Frequency Array optimization by highlighting the constraint (lowercase letters).
   - Discuss the language-specific trade-offs (Go's array keys vs. Java's String/List keys).

public List<List<String>> groupAnagrams(String[] strs) {
    if (strs == null || strs.length == 0) {
        return new ArrayList<>();
    }

    // Map to store the frequency signature as key and the list of anagrams as value
    Map<String, List<String>> groups = new HashMap<>();

    for (var s : strs) {
        // 1. Create frequency count for 'a'-'z'
		int[] count = new int[26];
		for (var c : s.toCharArray()) {
			count[c-'a']++;
		}

        // 2. Build a unique string key from the count array. Example: "eat" -> "#1#0#0#0#1...#1"
		var sb = new StringBuilder("");
		for (var d : count) {
			sb.append("#");
			sb.append(d);
		}
        var key = sb.toString();

        // 3. Group the original string
		if (!groups.containsKey(key)) {
			groups.put(key, new ArrayList<String>());
		}
		groups.get(key).add(s);
    }

    return new ArrayList<>(groups.values());
}
*/
