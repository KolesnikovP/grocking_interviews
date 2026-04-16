// Problem: Valid Anagram (#242)
// Platform: LeetCode
// Difficulty: Easy
// Date: 2026-04-15
// Time complexity: O(n)
// Space complexity: O(k) — k = character set size; O(1) for fixed alphabet
// Rating: PASS
// Notes: Frequency map with rune keys — increment for s, decrement for t, check all zeros.
//        Early exit on len(s) != len(t). Use m[val]++ not ok-check pattern.

package arrays

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, val := range s {
		m[val]++
	}

	for _, val := range t {
		m[val]--
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true
}
