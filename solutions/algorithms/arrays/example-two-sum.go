// Problem: Two Sum
// Platform: LeetCode
// Difficulty: Easy
// Date: YYYY-MM-DD
// Time complexity: O(n)
// Space complexity: O(n)
// Rating: PASS
// Notes: Hash map complement lookup — store value→index, check if (target - v) exists before inserting

// This file is an EXAMPLE showing the required solution format.
// Delete it and add your own solutions here.

package arrays

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		match := target - v
		if val, ok := m[match]; ok {
			return []int{i, val}
		}
		m[v] = i
	}

	return []int{}
}
