package domain

import "time"

type HabitLog struct {
	ID         int64
	HabitID    int64
	ExecutedAt time.Time
}

func NewHabitLog(habitID int64, executedAt time.Time) *HabitLog {
	return &HabitLog{HabitID: habitID, ExecutedAt: executedAt}
}
