# Contributing to grocking_interviews

Thanks for wanting to contribute. This is a lean project — the goal is to keep it useful, not to bloat it. Read this before opening a PR.

---

## What kind of contributions are welcome

- **New algorithm solutions** — more solved problems in any covered language (Go, JS, TS)
- **New solution directories** — if a topic folder doesn't exist yet (e.g., `linked-lists/`, `trees/`)
- **Agent improvements** — better hints, more realistic examples, improved scoring rubrics
- **New agent modes** — if a mode genuinely doesn't exist and solves a real gap
- **Bug fixes in agent logic** — wrong complexity analysis, broken routing, etc.
- **Curriculum additions** — new topics with a clear rationale

**Not welcome:**
- Cosmetic README changes
- Renaming things without a reason
- Adding dependencies (there are none and that's intentional)

---

## Adding a solution

1. Find the right directory under `solutions/algorithms/[topic]/`
   - If the directory doesn't exist, create it (e.g., `solutions/algorithms/linked-lists/`)
2. Name the file using the problem slug and language extension: `reverse-linked-list.go`
3. Include the standard header at the top:

```go
// Problem: Reverse Linked List
// Platform: LeetCode
// Difficulty: Easy
// Date: YYYY-MM-DD
// Time complexity: O(n)
// Space complexity: O(1)
// Rating: PASS
// Notes: iterative — track prev/curr/next pointers
```

4. Write clean, idiomatic code for the language
5. Open a PR with the title: `feat: add [problem name] ([language], [difficulty])`

### Supported languages

| Language | Extension | Style guide |
|---|---|---|
| Go | `.go` | Standard Go formatting — run `gofmt` before committing |
| JavaScript | `.js` | ES2020+, consistent with existing files |
| TypeScript | `.ts` | Strict mode, explicit types on function signatures |

---

## Adding or improving an agent

Agents live in `agents/` as markdown files. Each one is a prompt that Claude Code loads when the corresponding slash command is triggered.

### To improve an existing agent

1. Read the agent file fully before editing
2. Make targeted changes — don't rewrite the whole thing
3. Test it: open Claude Code in the repo, trigger the slash command, verify the behavior
4. In your PR description, explain specifically what was wrong and what you changed

### To add a new agent

Only add a new agent if there's a clear gap existing agents don't cover. Before writing one:

1. Open an issue explaining the use case
2. Describe what the agent does that existing agents don't
3. Wait for a thumbs-up before writing it

If approved, follow this structure:

```markdown
# Agent: [Name]

**Mode:** [MODE_NAME]
**Trigger:** `/command-name` — or [text trigger description]

---

## Your Role
[One paragraph — what the agent does and what it doesn't do]

---

## [Session flow or main logic sections]

---

## Output Format
[Exact format with example]

---

## Rules
[Hard constraints — things the agent must never do]
```

Then register it in `CLAUDE.md` (routing table + Agent Instructions section) and create the corresponding skill file in `.claude/skills/[command-name]/`.

---

## Modifying CLAUDE.md

`CLAUDE.md` is the system prompt. Changes here affect everything.

- Don't remove existing routing — it breaks backward compatibility
- Don't change progress tracking rules without updating `agents/organizer.md` to match
- Test any routing changes by verifying all slash commands still work after your edit

---

## PR process

1. Fork the repo
2. Create a branch: `feat/your-description` or `fix/your-description`
3. Make your changes
4. Test with Claude Code open in the repo
5. Open a PR with a clear title and description

**PR title format:**
- `feat: add [thing]`
- `fix: [what was broken]`
- `improve: [agent name] — [what changed]`
- `docs: [what was updated]`

**PR description should include:**
- What problem this solves
- How you tested it
- Any trade-offs or limitations

---

## Questions

Open an issue. Don't DM.
