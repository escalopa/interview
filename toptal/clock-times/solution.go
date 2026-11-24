package clocktimes

// CountValidTimes returns how many distinct valid 24-hour times HH:MM
// (00:00 .. 23:59) can be built using each of the four digits exactly once.
//
// There are only 4! = 24 orderings, so try them all. Duplicate digits create
// the same time more than once; a 1440-bit set (one bit per minute of the day,
// 23 uint64 words on the stack) removes duplicates without a map.
//
// Time O(4!) = O(1), space O(1).
func CountValidTimes(a, b, c, d int) int {
	digits := [4]int{a, b, c, d}
	var seen [23]uint64 // 23 * 64 = 1472 >= 1440 bits
	count := 0

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i {
				continue
			}
			hours := digits[i]*10 + digits[j]
			if hours > 23 {
				continue // prune: no minutes can fix the hour
			}
			for k := 0; k < 4; k++ {
				if k == i || k == j {
					continue
				}
				l := 6 - i - j - k // the remaining index (0+1+2+3 = 6)
				minutes := digits[k]*10 + digits[l]
				if minutes > 59 {
					continue
				}
				t := hours*60 + minutes
				word, bit := t/64, uint64(1)<<(t%64)
				if seen[word]&bit == 0 {
					seen[word] |= bit
					count++
				}
			}
		}
	}
	return count
}
