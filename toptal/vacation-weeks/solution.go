package vacationweeks

// VacationWeeks returns the number of full weeks in a vacation.
//
// Inputs:
//   - year: the calendar year (used only for the leap-year rule).
//   - startMonth, endMonth: months 1..12, inclusive, startMonth <= endMonth.
//   - jan1Weekday: weekday of 1 January, 0 = Monday ... 6 = Sunday.
//
// The vacation starts on the first Monday of startMonth and ends on the last
// Sunday of endMonth. The span Monday..Sunday is always a whole number of weeks.
//
// Work in "day of year" numbers (0 = 1 January). The weekday of day d is
// (jan1Weekday + d) % 7. Only the 12 month lengths are summed, so
// time O(1) (at most 12 steps) and space O(1).
func VacationWeeks(year, startMonth, endMonth, jan1Weekday int) int {
	monthDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeap(year) {
		monthDays[1] = 29
	}

	// Day of year of the 1st of startMonth, and of the last day of endMonth.
	firstDay, lastDay := 0, 0
	for m := 0; m < endMonth; m++ {
		if m < startMonth-1 {
			firstDay += monthDays[m]
		}
		lastDay += monthDays[m]
	}
	lastDay-- // day after the end of endMonth, minus one

	// Move forward to the first Monday (weekday 0).
	wd := (jan1Weekday + firstDay) % 7
	start := firstDay + (7-wd)%7

	// Move back to the last Sunday (weekday 6).
	wd = (jan1Weekday + lastDay) % 7
	end := lastDay - (wd+1)%7

	if end < start {
		return 0
	}
	return (end - start + 1) / 7
}

func isLeap(y int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}
