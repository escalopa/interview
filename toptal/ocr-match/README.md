# OCR match

## Problem

Two strings come from OCR (text recognition). Each string contains:

- letters `a-z` and `A-Z`: characters that were recognized;
- digits `0-9`: a run of digits is a number, the count of characters that
  could not be recognized.

Return `true` if both strings can describe the same original text, else `false`.

### Examples

| a      | b        | result | why                                    |
|--------|----------|--------|----------------------------------------|
| `A2Le` | `2pL1`   | `true` | `A??Le` vs `??pL?`: both have 5 chars, no conflict |
| `a10`  | `10a`    | `true` | both are 11 chars, `a` meets an unknown |
| `ba1`  | `1Ad`    | `false`| `a` vs `A` at index 1                  |
| `3x2`  | `abcyyz` | `false`| `x` vs `y` at index 3                  |
| `a12b` | `a12`    | `false`| lengths 14 vs 13                       |

## Clarifying questions

- Is `12` one number (12 unknown chars) or two (1 + 2)?
  **Answer: one number.** If any split were allowed, the task becomes
  LeetCode 2060 and needs DP.
- Are letters case-sensitive? **Assumed: yes.**
- Can `0` appear? **Yes, it means zero unknown chars.**
- Can numbers be huge? **Assumed: they fit in `int`.**

## Solution

Two pointers, one per string. Each side keeps a counter of pending unknown chars.

1. If a side has no pending unknowns and stands on a digit, read the whole number.
2. Both sides have unknowns: subtract the smaller counter from both in one step.
   So `a1000000b` costs O(1), not one million steps.
3. Only one side has unknowns: one unknown eats one letter from the other side.
4. Letter vs letter: they must be equal.
5. At the end, both strings must be fully used and both counters must be `0`.

The expanded strings are never built.

**Complexity:** time `O(n + m)`, space `O(1)`.

Code: [solution.go](solution.go). Tests compare against a brute force that
expands both strings, on 200 000 random pairs: [solution_test.go](solution_test.go).
