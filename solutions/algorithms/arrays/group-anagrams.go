// Problem: Group Anagrams (#49)
// Platform: LeetCode
// Difficulty: Medium
// Date: 2026-04-16
// Time complexity: O(n * k log k) — n words, each sorted in O(k log k)
// Space complexity: O(n * k) — map stores all words
// Rating: PASS
// Notes: Sort each word to get canonical key → group into map[string][]string.
//        append(m[key], word) works without ok-check — nil slice appends cleanly.
//        camelCase: wordChars not word_chars (Go convention).

package arrays

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, word := range strs {
		wordChars := []rune(word)
		sort.Slice(wordChars, func(i, j int) bool {
			return wordChars[i] < wordChars[j]
		})

		key := string(wordChars)
		m[key] = append(m[key], word)
	}

	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}

	return result
}
