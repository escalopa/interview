package clocktimes

import "testing"

// bruteTimes checks every minute of the day against the digit multiset.
func bruteTimes(a, b, c, d int) int {
	var want [10]int
	for _, x := range []int{a, b, c, d} {
		want[x]++
	}
	count := 0
	for h := range 24 {
		for m := range 60 {
			var got [10]int
			got[h/10]++
			got[h%10]++
			got[m/10]++
			got[m%10]++
			if got == want {
				count++
			}
		}
	}
	return count
}

func TestCountValidTimes(t *testing.T) {
	cases := []struct {
		a, b, c, d, want int
	}{
		{1, 8, 3, 2, 6},
		{2, 3, 3, 2, 3},
		{6, 2, 4, 7, 0},
		{0, 0, 0, 0, 1},
	}
	for _, c := range cases {
		if got := CountValidTimes(c.a, c.b, c.c, c.d); got != c.want {
			t.Errorf("%d,%d,%d,%d: got %d, want %d", c.a, c.b, c.c, c.d, got, c.want)
		}
	}
}

func TestCountValidTimesExhaustive(t *testing.T) {
	for x := range 10000 {
		a, b, c, d := x/1000, x/100%10, x/10%10, x%10
		if got, want := CountValidTimes(a, b, c, d), bruteTimes(a, b, c, d); got != want {
			t.Fatalf("%d,%d,%d,%d: got %d, want %d", a, b, c, d, got, want)
		}
	}
}
