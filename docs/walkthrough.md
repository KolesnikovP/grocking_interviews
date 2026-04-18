# Full workflow walkthrough

This is a complete session from scratch — picking a problem, getting stuck, working through hints, reviewing the solution, and logging progress. Every command is real. Every response is representative of what you'll actually see.

The problem used here is **Two Sum** (LeetCode #1, Easy). It's the first problem on LeetCode for a reason — the brute force solution is obvious, and the optimal solution introduces one of the most important patterns you'll use everywhere: the hash map lookup.

---

## 1. Open the repo in Claude Code

```bash
git clone https://github.com/KolesnikovP/grocking_interviews.git
cd grocking_interviews
claude
```

Claude Code loads `CLAUDE.md` automatically. All agents and slash commands are available immediately.

---

## 2. Ask what to study next

If you're not sure where to start, ask the organizer:

> **You:** `/gi-organize` what should I work on next?

> **Agent:**
> ```
> [ORGANIZE MODE]
>
> Your tracker is empty — nothing practiced yet.
>
> Recommended starting point: Arrays & Strings.
> Good first problem: Two Sum (LeetCode #1)
> Pattern: Hash Map
> Difficulty: Easy
> ```

> 💡 The organizer reads `progress/tracker.md` to make this recommendation. On a fresh install, it points you to the beginning of the curriculum.

---

## 3. Attempt the problem

Go to [LeetCode #1](https://leetcode.com/problems/two-sum/). Read it. Give yourself 10–15 minutes before asking for help.

**The problem:** Given an array of integers and a target, return the indices of the two numbers that add up to the target.

A natural first attempt — check every pair:

```go
// first attempt — O(n²)
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}
```

It works. But there's a faster way. Before submitting, get it reviewed.

---

## 4. Run the reviewer

> **You:** `/gi-review`
>
> Problem: Two Sum (LeetCode #1)
> Language: Go
>
> ```go
> func twoSum(nums []int, target int) []int {
>     for i := 0; i < len(nums); i++ {
>         for j := i + 1; j < len(nums); j++ {
>             if nums[i]+nums[j] == target {
>                 return []int{i, j}
>             }
>         }
>     }
>     return []int{}
> }
> ```

> **Agent:**
> ```
> [REVIEW MODE]
>
> Pattern: Hash Map — this is the classic complement lookup problem.
> You applied brute force instead. That's the main finding.
>
> Correctness: ✓ — produces correct output for all standard cases.
>
> Edge cases missed:
> - Empty array (would return [] — acceptable)
> - No valid pair exists (problem guarantees one exists, so this is fine)
>
> Time complexity:  O(n²) — every pair checked once
> Space complexity: O(1) — no auxiliary structures
>
> Code quality:
> - Clean and readable
> - Logic is easy to follow
>
> Rating: BORDERLINE
> Correct, but O(n²) won't pass on large inputs. A linear solution
> exists using a hash map — expected in interviews.
>
> Type `/gi-improve` to see the optimized version.
> ```

> 💡 BORDERLINE means: correct logic, missed a more efficient pattern. The reviewer never shows you the better solution — that's what `/gi-improve` is for.

---

## 5. Optimize it

> **You:** `/gi-improve`

> **Agent:**
> ```
> [OPTIMIZE MODE]
>
> Optimal pattern: Hash Map (complement lookup)
>
> The key insight: instead of checking every pair, ask a different
> question at each step. For the current number, what number would
> complete the target? Look that number up instantly in a map
> instead of scanning the rest of the array.
>
> Bottleneck: the nested loop — O(n²) because for every element
> you scan all remaining elements.
>
> Complexity comparison:
> |        | Time  | Space |
> |--------|-------|-------|
> | Before | O(n²) | O(1)  |
> | After  | O(n)  | O(n)  |
>
> Optimized solution:
>
> func twoSum(nums []int, target int) []int {
>     m := make(map[int]int)
>
>     for i, v := range nums {
>         match := target - v
>         if val, ok := m[match]; ok {
>             return []int{i, val}
>         }
>         m[v] = i
>     }
>
>     return []int{}
> }
>
> Why this works: for each number, compute its complement
> (target - v). Check if that complement is already in the map.
> If yes — you found your pair. If no — store the current number
> and move on. One pass, done.
>
> Note: we trade O(1) space for O(n) space to get from O(n²) to O(n)
> time. In interviews, this trade is almost always correct.
> ```

> 💡 The agent always explains the trade-off (time vs space). Understanding why a solution is better — not just that it is — is what transfers to new problems.

---

## 6. Write the final solution file

```bash
touch solutions/algorithms/arrays/two-sum.go
```

```go
// Problem: Two Sum
// Platform: LeetCode
// Difficulty: Easy
// Date: 2026-04-18
// Time complexity: O(n)
// Space complexity: O(n)
// Rating: PASS (after /gi-improve from BORDERLINE)
// Notes: Hash map complement lookup — store value→index, check if (target - v) exists before inserting

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
```

---

## 7. Log it with the organizer

> **You:** `/gi-organize`
>
> Problem: Two Sum
> Platform: LeetCode
> Topic: Arrays / Hash Tables
> Difficulty: Easy
> Language: Go
> Reviewer rating: PASS (after improve)
> Stuck: no
> Notes: Started with brute force O(n²), reviewer flagged BORDERLINE, improved to O(n) with hash map complement lookup

> **Agent:**
> ```
> [ORGANIZE MODE]
>
> Filed: solutions/algorithms/arrays/two-sum.go
> Tracker updated: Arrays & Strings → In Progress (1 session)
>                  Hash Tables → In Progress (1 session)
>
> Progress snapshot:
> - Confident:    0 topics
> - In Progress:  2 topics (Arrays & Strings, Hash Tables)
> - Not Started: 37 topics
>
> Recommended next: stay on Arrays — try Contains Duplicate (LeetCode #217)
> Pattern: Hash Set
> Difficulty: Easy
> ```

---

## 8. Commit your solution

```bash
git add solutions/algorithms/arrays/two-sum.go
git add progress/tracker.md
git commit -m "feat: add Two Sum (Go, Easy, PASS)"
git push
```

> 💡 Commit to your **private** personal repo, not this public template. See the [two-repo setup guide](../SETUP_TWO_REPOS.md) if you haven't done that yet.

---

## What you just did

| Step | Command | What happened |
|---|---|---|
| Picked a problem | `/gi-organize` | Organizer read the tracker and recommended the first problem |
| Got a review | `/gi-review` | Rated BORDERLINE — correct but O(n²) |
| Optimized | `/gi-improve` | Agent explained the hash map complement lookup |
| Logged progress | `/gi-organize` | Tracker updated, next problem recommended |

---

## What if you were actually stuck?

If you couldn't write anything at all, use `/gi-teacher` before reviewing:

> **You:** `/gi-teacher`
> I'm working on Two Sum. I don't know where to start.

> **Agent:**
> ```
> [STUCK MODE — Socratic]
>
> Let's start simple. What do you know about the two numbers
> you're looking for?
>
> If one of them is 3 and the target is 9, what must the other one be?
> ```

The teacher asks questions until you reach the "complement lookup" insight yourself. It won't show you the solution — that's the point. Once you work it out, proceed to `/gi-review` as normal.

---

## Full command reference

| Command | When to use |
|---|---|
| `/gi-teacher` | Before writing any solution — use when stuck or unsure of approach |
| `/gi-review` | After writing a solution — correctness, complexity, rating |
| `/gi-improve` | After a BORDERLINE rating — get the optimal version with explanation |
| `/gi-design` | For system design problems — full session with scoring |
| `/gi-pattern` | To learn or practice a specific pattern or OOP concept |
| `/gi-behavioral` | To build and drill STAR stories for behavioral interviews |
| `/gi-organize` | After any solved problem — files solution, updates tracker, recommends next |
