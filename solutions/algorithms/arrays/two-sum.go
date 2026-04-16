// Problem: Two Sum
// Platform: LeetCode
// Difficulty: Easy
// Date: 2026-04-12
// Time complexity: O(n)
// Space complexity: O(n)
// Rating: PASS
// Notes: Hash map complement lookup — store value→index, check if (target - v) exists before inserting

package arrays

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		match := target - v
		val, ok := m[match]
		if ok {
			return []int{i, val}
		} else {
			m[v] = i
		}
	}

	return []int{}
}
