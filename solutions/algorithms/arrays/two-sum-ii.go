// Problem: Two Sum II - Input Array Is Sorted (#167)
// Platform: LeetCode
// Difficulty: Medium
// Date: 2026-04-17
// Time complexity: O(n)
// Space complexity: O(1)
// Rating: PASS
// Notes: Two Pointers pattern — sorted array allows narrowing from both ends.
//        sum > target → right--, sum < target → left++. Return 1-indexed.
//        No hash map needed — O(1) space constraint forces Two Pointers.

package arrays

func twoSumII(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
