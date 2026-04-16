// Problem: Contains Duplicate (#217)
// Platform: LeetCode
// Difficulty: Easy
// Date: 2026-04-13
// Time complexity: O(n)
// Space complexity: O(n)
// Rating: PASS (after optimize)
// Notes: Use map[int]struct{} (set) — not map[int]int. struct{} costs 0 bytes; storing index is wasteful when existence is all you need.

package arrays

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}

	return false
}
