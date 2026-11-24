package ocrmatch

// CanMatch reports whether two OCR strings can describe the same original text.
//
// Letters (a-z, A-Z) are known characters and are compared case-sensitively.
// A run of digits is one decimal number: the count of unknown characters
// ("a12b" = 'a', 12 unknown characters, 'b'). An unknown character matches anything.
//
// Two pointers walk both strings at once. Each side keeps a "pending unknowns"
// counter. When both sides have pending unknowns, the smaller counter is
// subtracted from both in one step, so large numbers cost O(1), not O(number).
//
// Time O(len(a) + len(b)), space O(1).
// Assumes each number fits in int (true for any realistic input).
func CanMatch(a, b string) bool {
	i, j := 0, 0         // positions in a and b
	skipA, skipB := 0, 0 // pending unknown characters on each side

	for {
		// Load the next number on a side that has no pending unknowns.
		// The loop also skips zero counts such as "a0b".
		for skipA == 0 && i < len(a) && isDigit(a[i]) {
			skipA, i = readNumber(a, i)
		}
		for skipB == 0 && j < len(b) && isDigit(b[j]) {
			skipB, j = readNumber(b, j)
		}

		switch {
		case skipA > 0 && skipB > 0:
			// Unknown vs unknown: consume the overlap in one step.
			k := min(skipA, skipB)
			skipA -= k
			skipB -= k
		case skipA > 0:
			// Unknown in a vs known letter in b.
			if j == len(b) {
				return false
			}
			skipA--
			j++
		case skipB > 0:
			// Known letter in a vs unknown in b.
			if i == len(a) {
				return false
			}
			skipB--
			i++
		default:
			// Letter vs letter, or end of input.
			if i == len(a) || j == len(b) {
				return i == len(a) && j == len(b)
			}
			if a[i] != b[j] {
				return false
			}
			i++
			j++
		}
	}
}

func isDigit(c byte) bool { return c >= '0' && c <= '9' }

// readNumber parses the digit run that starts at s[i].
// It returns the number and the index after the run.
func readNumber(s string, i int) (int, int) {
	n := 0
	for i < len(s) && isDigit(s[i]) {
		n = n*10 + int(s[i]-'0')
		i++
	}
	return n, i
}
