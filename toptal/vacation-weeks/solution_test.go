package vacationweeks

import (
	"testing"
	"time"
)

func TestVacationWeeksExample(t *testing.T) {
	// 2014, April..May, 1 Jan 2014 was a Wednesday (2). Expected 7 weeks.
	if got := VacationWeeks(2014, 4, 5, 2); got != 7 {
		t.Fatalf("got %d, want 7", got)
	}
}

// bruteWeeks walks real calendar days with the time package.
func bruteWeeks(year, a, b int) int {
	d := time.Date(year, time.Month(a), 1, 0, 0, 0, 0, time.UTC)
	for d.Weekday() != time.Monday {
		d = d.AddDate(0, 0, 1)
	}
	end := time.Date(year, time.Month(b)+1, 0, 0, 0, 0, 0, time.UTC) // last day of b
	for end.Weekday() != time.Sunday {
		end = end.AddDate(0, 0, -1)
	}
	if end.Before(d) {
		return 0
	}
	return (int(end.Sub(d).Hours()/24) + 1) / 7
}

func TestVacationWeeksAgainstCalendar(t *testing.T) {
	for y := 1900; y <= 2100; y++ {
		jan1 := (int(time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC).Weekday()) + 6) % 7 // Monday = 0
		for a := 1; a <= 12; a++ {
			for b := a; b <= 12; b++ {
				if got, want := VacationWeeks(y, a, b, jan1), bruteWeeks(y, a, b); got != want {
					t.Fatalf("year %d months %d..%d: got %d, want %d", y, a, b, got, want)
				}
			}
		}
	}
}
