# Full workflow walkthrough

This is a complete session from scratch — picking a problem, getting stuck, working through hints, reviewing the solution, and logging progress. Every command is real. Every response is representative of what you'll actually see.

The problem used here is **Best Time to Buy and Sell Stock** (LeetCode #121, Easy). It's a good first example because most people's instinct is a brute-force O(n²) approach, and the optimal solution requires a small but non-obvious insight.

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
> Based on your tracker:
> - Arrays & Strings: Practiced (5 sessions)
> - Hash Tables: Practiced (4 sessions)
> - Everything else: Not Started
>
> Recommended next: continue Arrays — you haven't hit Sliding Window yet.
> Good next problem: Best Time to Buy and Sell Stock (LeetCode #121)
> Pattern: Sliding Window / one-pass scan
> ```

> 💡 The organizer reads `progress/tracker.md` to make this recommendation. It always prioritizes in-progress topics over untouched ones.

---

## 3. Attempt the problem

Go to [LeetCode #121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/). Read it. Give yourself 10–15 minutes before asking for help.

A common first attempt — check every pair of days:

```go
// first attempt — O(n²)
func maxProfit(prices []int) int {
    maxP := 0
    for i := 0; i < len(prices); i++ {
        for j := i + 1; j < len(prices); j++ {
            profit := prices[j] - prices[i]
            if profit > maxP {
                maxP = profit
            }
        }
    }
    return maxP
}
```

It's correct. But it's slow. Before submitting, ask for a review.

---

## 4. Run the reviewer

> **You:** `/gi-review`
>
> Problem: Best Time to Buy and Sell Stock (LeetCode #121)
> Language: Go
>
> ```go
> func maxProfit(prices []int) int {
>     maxP := 0
>     for i := 0; i < len(prices); i++ {
>         for j := i + 1; j < len(prices); j++ {
>             profit := prices[j] - prices[i]
>             if profit > maxP {
>                 maxP = profit
>             }
>         }
>     }
>     return maxP
> }
> ```

> **Agent:**
> ```
> [REVIEW MODE]
>
> Pattern: Sliding Window — this is a classic one-pass scan.
> You applied brute force instead. That's the main finding.
>
> Correctness: ✓ — produces correct output for all standard cases.
>
> Edge cases missed:
> - Single-element array (returns 0 — correct, but worth verifying)
> - Already sorted descending (returns 0 — correct)
>
> Time complexity:  O(n²) — nested loops, each pair evaluated once
> Space complexity: O(1) — no auxiliary structures
>
> Code quality:
> - Variable names are clear (maxP, profit)
> - Logic is easy to follow
>
> Rating: BORDERLINE
> Correct, but the nested loop won't pass on large inputs (n = 100,000).
> A linear solution exists and is expected here.
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
> Optimal pattern: Sliding Window (one-pass scan)
>
> The key insight: you don't need to check every pair.
> As you walk forward, track the lowest price seen so far.
> At each step, the best profit you can make TODAY is prices[i] - minSoFar.
>
> Bottleneck: the nested loop — O(n²) because you re-check
> every previous day for every current day.
>
> Complexity comparison:
> |        | Time  | Space |
> |--------|-------|-------|
> | Before | O(n²) | O(1)  |
> | After  | O(n)  | O(1)  |
>
> Optimized solution:
>
> func maxProfit(prices []int) int {
>     minPrice := prices[0]
>     maxP := 0
>
>     for _, price := range prices {
>         if price < minPrice {
>             minPrice = price
>         } else if price-minPrice > maxP {
>             maxP = price - minPrice
>         }
>     }
>
>     return maxP
> }
>
> Why this works: you're sliding a conceptual window forward.
> The left edge is always the cheapest day seen so far.
> The right edge is today. You never need to look back further
> than your current minimum — if a cheaper day existed before it,
> that IS your minimum already.
> ```

> 💡 The agent always shows before/after complexity. That comparison is the point — understanding *what you traded* is how pattern recognition builds.

---

## 6. Write the final solution file

```bash
touch solutions/algorithms/arrays/best-time-to-buy-and-sell-stock.go
```

```go
// Problem: Best Time to Buy and Sell Stock
// Platform: LeetCode
// Difficulty: Easy
// Date: 2026-04-18
// Time complexity: O(n)
// Space complexity: O(1)
// Rating: PASS (after /gi-improve from BORDERLINE)
// Notes: Sliding window — track minPrice as left edge, compute profit at each step

package arrays

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxP := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxP {
			maxP = price - minPrice
		}
	}

	return maxP
}
```

---

## 7. Log it with the organizer

> **You:** `/gi-organize`
>
> Problem: Best Time to Buy and Sell Stock
> Platform: LeetCode
> Topic: Arrays / Sliding Window
> Difficulty: Easy
> Language: Go
> Reviewer rating: PASS (after improve)
> Stuck: no
> Notes: Started brute force O(n²), reviewer flagged it, improved to O(n) with sliding window

> **Agent:**
> ```
> [ORGANIZE MODE]
>
> Filed: solutions/algorithms/arrays/best-time-to-buy-and-sell-stock.go
> Tracker updated: Arrays & Strings → Practiced (6 sessions)
>
> Progress snapshot:
> - Confident:    0 topics
> - Practiced:    2 topics (Arrays & Strings, Hash Tables)
> - Not Started: 37 topics
>
> Recommended next: continue Arrays — try a Medium problem next
> to push toward Confident status.
> Good next problem: Container With Most Water (LeetCode #11)
> Pattern: Two Pointers
> ```

---

## 8. Commit your solution

```bash
git add solutions/algorithms/arrays/best-time-to-buy-and-sell-stock.go
git add progress/tracker.md
git commit -m "feat: add Best Time to Buy and Sell Stock (Go, Easy, PASS)"
git push
```

> 💡 Commit to your **private** personal repo, not this public template. See the [two-repo setup guide](../SETUP_TWO_REPOS.md) if you haven't done that yet.

---

## What you just did

| Step | Command | What happened |
|---|---|---|
| Picked a problem | `/gi-organize` | Organizer read the tracker and surfaced the next gap |
| Got a review | `/gi-review` | Rated BORDERLINE — correct but O(n²) |
| Optimized | `/gi-improve` | Agent explained the Sliding Window insight |
| Logged progress | `/gi-organize` | Tracker updated, next problem recommended |

---

## What if you were actually stuck?

If you couldn't write anything at all, use `/gi-teacher` before reviewing:

> **You:** `/gi-teacher`
> I'm working on Best Time to Buy and Sell Stock.
> I don't know how to approach it — checking every pair feels wrong.

> **Agent:**
> ```
> [STUCK MODE — Socratic]
>
> Good instinct. Let's think about what you actually need.
>
> If you're standing on day 5 and thinking about selling,
> what's the only thing that determines your profit?
> ```

The teacher asks questions until you reach the "track the minimum" insight yourself. It won't show you the solution — that's the point. Once you work it out, proceed to `/gi-review` as normal.

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
