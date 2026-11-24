package ocrmatch

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestCanMatch(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"A2Le", "2pL1", true},
		{"a10", "10a", true},
		{"ba1", "1Ad", false},
		{"3x2", "abcxyz", true},  // ???x?? vs abcxyz
		{"3x2", "abcyyz", false}, // 'x' vs 'y' at index 3
		{"3x2", "abcx2", true},
		{"a12b", "a12b", true},
		{"a12b", "a13", true},
		{"a12b", "a12", false}, // lengths 14 vs 13
		{"a0b", "ab", true},
		{"", "", true},
		{"", "0", true},
		{"", "1", false},
		{"A", "a", false}, // case-sensitive
		{"100", "99z", true},
	}
	for _, c := range cases {
		if got := CanMatch(c.a, c.b); got != c.want {
			t.Errorf("CanMatch(%q, %q) = %v, want %v", c.a, c.b, got, c.want)
		}
		if got := CanMatch(c.b, c.a); got != c.want {
			t.Errorf("CanMatch(%q, %q) = %v, want %v (swapped)", c.b, c.a, got, c.want)
		}
	}
}

// expand turns "a3b" into "a???b" for the brute-force reference.
func expand(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); {
		if isDigit(s[i]) {
			n, k := readNumber(s, i)
			sb.WriteString(strings.Repeat("?", n))
			i = k
		} else {
			sb.WriteByte(s[i])
			i++
		}
	}
	return sb.String()
}

func bruteMatch(a, b string) bool {
	x, y := expand(a), expand(b)
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != '?' && y[i] != '?' && x[i] != y[i] {
			return false
		}
	}
	return true
}

func randomOCR(r *rand.Rand) string {
	var sb strings.Builder
	for n := r.Intn(6); n > 0; n-- {
		if r.Intn(2) == 0 {
			sb.WriteString(strconv.Itoa(r.Intn(15)))
		} else {
			sb.WriteByte("abAB"[r.Intn(4)])
		}
	}
	return sb.String()
}

func TestCanMatchRandom(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for range 200000 {
		a, b := randomOCR(r), randomOCR(r)
		if got, want := CanMatch(a, b), bruteMatch(a, b); got != want {
			t.Fatalf("CanMatch(%q, %q) = %v, want %v", a, b, got, want)
		}
	}
}
