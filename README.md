# grocking_interviews

A personal interview prep system built on top of [Claude Code](https://claude.ai/code). Instead of a chatbot you talk to, it's a set of specialized AI agents that each do one job — review your code, teach when you're stuck, coach system design, prep behavioral answers, and track everything.

Everything runs inside your editor. No web app, no subscriptions beyond Claude Code itself.

→ **[Full workflow walkthrough](docs/walkthrough.md)** — see a complete session from picking a problem to logging it in the tracker.

---

## What's inside

Seven agents, each triggered by a slash command:

| Command | What it does |
|---|---|
| `/gi-review` | Senior engineer code review — correctness, complexity, code quality, pass/fail rating |
| `/gi-teacher` | Stuck? Progressive hints + Socratic questions. Never gives the answer outright |
| `/gi-improve` | Takes a working solution and makes it faster/cleaner. Explains why |
| `/gi-design` | Staff engineer running a system design session — requirements, HLD, deep dive, scoring |
| `/gi-pattern` | Explains OOP/SOLID/GoF patterns with real examples. Has a quiz mode |
| `/gi-behavioral` | STAR framework coach — builds your story bank, drills you on delivery |
| `/gi-organize` | Files your solution, updates the progress tracker, recommends what to study next |

---

## Prerequisites

- [Claude Code](https://claude.ai/code) installed and authenticated
- Go installed (for algorithm solutions) — or JavaScript/TypeScript if you prefer
- Git

---

## Setup

```bash
git clone https://github.com/KolesnikovP/grocking_interviews.git
cd grocking_interviews
claude
```

That's it. Claude Code picks up `CLAUDE.md` automatically and loads all agents.

---

## How to use

### Solving an algorithm problem

1. Write your solution in `solutions/algorithms/[topic]/[problem-name].go`
2. Open Claude Code in the repo
3. Paste your solution and type `/gi-review`

The reviewer will check correctness, identify the pattern, give you time/space complexity, and rate it PASS / BORDERLINE / NEEDS WORK.

```
/gi-review

Problem: Two Sum (LeetCode #1)
Language: Go

func twoSum(nums []int, target int) []int {
    // your solution here
}
```

### When you're stuck

Type `/gi-teacher` before submitting any code. The teacher won't give you the answer — it'll ask you questions until you figure it out. If you're completely blocked after three rounds of hints, you can explicitly ask for the full solution.

```
/gi-teacher

I'm working on Group Anagrams. I know I need to group strings that are anagrams
of each other but I don't know how to identify which strings belong together.
```

### Optimizing a working solution

Already have a PASS but know your solution is O(n²)? Use `/gi-improve`:

```
/gi-improve

Language: Go
[paste your working solution]
```

### System design practice

```
/gi-design

Design a URL shortener
```

The agent will run you through requirements gathering, HLD, a deep dive on the hardest component, and score your design out of 15.

### OOP and design patterns

```
/gi-pattern

Explain the Strategy pattern

# or quiz mode:
/gi-pattern

Quiz me on the Adapter pattern
```

### Behavioral prep

```
/gi-behavioral

Let's build a conflict story from when I pushed back on a product decision at my last job
```

### Logging progress

After any PASS-rated solution, type `/gi-organize` to file it properly and update the tracker. The organizer also tells you what to study next based on your weakest areas.

---

## Project structure

```
grocking_interviews/
├── CLAUDE.md                      # System prompt — loads all agents and routing logic
├── agents/
│   ├── reviewer.md                # /gi-review agent
│   ├── teacher.md                 # /gi-teacher agent
│   ├── improver.md                # /gi-improve agent
│   ├── system-design.md           # /gi-design agent
│   ├── oop-patterns.md            # /gi-pattern agent
│   ├── behavioral.md              # /gi-behavioral agent
│   └── organizer.md               # /gi-organize agent
├── docs/
│   └── walkthrough.md             # Full end-to-end workflow example
├── progress/
│   ├── tracker.md                 # Single source of truth for learning progress
│   └── story-bank.md              # Behavioral story bank (STAR format)
└── solutions/
    └── algorithms/
        └── arrays/                # Solutions with metadata headers
            └── example-two-sum.go
```

---

## Curriculum

The system covers everything for a mid-to-senior SWE interview:

**1. Data Structures & Algorithms**
Arrays, Linked Lists, Stacks & Queues, Trees & Graphs, Hash Tables, Sorting & Binary Search

**1b. Algorithmic Patterns**
Sliding Window, Two Pointers, Fast & Slow Pointers, Merge Intervals, Cyclic Sort, In-place Reversal, Tree BFS/DFS, Two Heaps, Subsets, Modified Binary Search, Top K Elements, K-way Merge, Topological Sort, Backtracking, Dynamic Programming

**2. System Design (HLD)**
Scalability, CAP theorem, Databases, Caching, Load Balancing, Message Queues, Storage — full design sessions for Twitter, URL Shortener, Chat App, YouTube, Uber

**3. Object-Oriented Design (LLD)**
SOLID principles, all 23 GoF patterns (Creational, Structural, Behavioral), Clean Code

**4. Language Mastery**
JavaScript (closures, event loop, async/await), TypeScript (generics, utility types), Go (goroutines, channels, interfaces)

**5. React / React Native**
Hooks, navigation, state management (Redux, Zustand), performance optimization

**6. Behavioral (STAR method)**
Conflict, failure, leadership, impact, ambiguity, collaboration, growth

---

## Progress tracking

`progress/tracker.md` is updated automatically by `/gi-organize` after every session. It tracks status per topic (Not Started → In Progress → Practiced → Confident), session counts, key insights, and a full log of solved problems. The organizer uses it to recommend what to study next, prioritizing weak areas.

---

## Solution file format

```go
// Problem: [title]
// Platform: [LeetCode / Codewars / HackerRank / other]
// Difficulty: [Easy / Medium / Hard]
// Date: [YYYY-MM-DD]
// Time complexity: O(...)
// Space complexity: O(...)
// Rating: [PASS / BORDERLINE / NEEDS WORK]
// Notes: [key insight in one line]
```

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

---

## License

MIT
