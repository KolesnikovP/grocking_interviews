// Problem: Find the Smallest Integer in an Array
// Platform: Codewars
// Difficulty: Easy
// Date: 2026-04-10
// Time complexity: O(n)
// Space complexity: O(1)
// Rating: BORDERLINE
// Notes: Seed value from numbers[0] is idiomatic — no need for math.MaxInt. Added empty slice guard.

package kata

func SmallestIntegerFinder(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	smallest := numbers[0]
	for _, v := range numbers {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}
