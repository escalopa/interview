# Vacation weeks

## Problem

Given a year `Y`, a start month `A`, an end month `B`, and the weekday `W` of
1 January of that year, return the number of full weeks of vacation.

- The vacation can start only on a Monday: the first Monday of month `A`.
- The vacation can end only on a Sunday: the last Sunday of month `B`.

### Example

`Y = 2014`, `A = April`, `B = May`, `W = Wednesday` → `7`.

First Monday of April 2014 is 7 April. Last Sunday of May 2014 is 25 May.
7 April .. 25 May is 49 days = 7 weeks.

## Clarifying questions

- Is the weekday of 1 January given? **Answer: yes, as input `W`.**
- Are both months included? **Assumed: yes.**
- Can the vacation cross into the next year? **Assumed: no, `A <= B`.**
- Leap years? **Yes: divisible by 4, except centuries not divisible by 400.**
- Input format: **months `1..12`, weekday `0 = Monday .. 6 = Sunday`.**

## Solution

Number the days of the year from `0` (1 January). Day `d` falls on weekday
`(W + d) % 7`.

1. `firstDay` = day number of the 1st of `A` (sum of month lengths before `A`).
2. `lastDay` = day number of the last day of `B`.
3. Move forward to the first Monday: `start = firstDay + (7 - wd) % 7`.
4. Move back to the last Sunday: `end = lastDay - (wd + 1) % 7`.
5. Answer: `(end - start + 1) / 7`. It divides exactly, because the span runs
   Monday .. Sunday. If `end < start`, the answer is `0`.

**Complexity:** time `O(1)` (at most 12 additions), space `O(1)`.

Code: [solution.go](solution.go). Tests compare against a day-by-day walk with
the `time` package for every year 1900..2100 and every month range:
[solution_test.go](solution_test.go).
