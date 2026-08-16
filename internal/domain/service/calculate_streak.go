package service

import (
	"time"
)

func CalculateStreak(executionsDate []time.Time, now time.Time) int64 {
	if len(executionsDate) == 0 {
		return 0
	}
	lastExecutionDateIdx := len(executionsDate) - 1
	if isStreakBroken(simplifyDate(executionsDate[lastExecutionDateIdx]), simplifyDate(now)) {
		return 0
	}
	streak := int64(1)
	for i := lastExecutionDateIdx; i > 0; i-- {
		current := simplifyDate(executionsDate[i])
		previous := simplifyDate(executionsDate[i-1])
		if current.Equal(previous) {
			continue
		}
		if isADayAfter(current, previous) {
			streak++
			continue
		}
		return streak
	}
	return streak
}

func isStreakBroken(t time.Time, now time.Time) bool {
	yesterday := now.AddDate(0, 0, -1)
	if t.Equal(now) {
		return false
	}
	if t.Equal(yesterday) {
		return false
	}
	return true
}

// returns true if t1 is a day after t2
func isADayAfter(t1, t2 time.Time) bool {
	return t1.Equal(t2.AddDate(0, 0, 1))
}

func simplifyDate(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
