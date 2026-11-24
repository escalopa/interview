# Valid clock times

## Problem

Given four digits, count how many different valid 24-hour times `HH:MM` can be
built. Each digit must be used exactly once. Valid times are `00:00 .. 23:59`.

### Examples

| digits     | result | times                                         |
|------------|--------|-----------------------------------------------|
| `1 8 3 2`  | `6`    | 12:38, 13:28, 18:23, 18:32, 21:38, 23:18      |
| `2 3 3 2`  | `3`    | 22:33, 23:23, 23:32                           |
| `6 2 4 7`  | `0`    | no valid time                                 |
| `0 0 0 0`  | `1`    | 00:00                                         |

## Clarifying questions

- Count, list, or the latest time? **Answer: count.**
- Same time from repeated digits counted once? **Yes, count distinct times.**
- Must all 4 digits be used? **Yes.**

## Solution

There are only `4! = 24` orders of the digits, so try them all.

1. Pick the two hour digits. If the hour is over `23`, skip: no minutes can fix it.
2. The last two digits are the minutes. Skip if over `59`.
3. Repeated digits make the same time more than once. A 1440-bit set (one bit
   per minute of the day, 23 `uint64` words on the stack, no map) counts each
   time once.

**Complexity:** time `O(1)` (24 orders), space `O(1)`.

Code: [solution.go](solution.go). Tests compare against a brute force over all
1440 minutes, for all 10 000 digit inputs: [solution_test.go](solution_test.go).
