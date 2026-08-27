package service

import (
	"testing"
	"time"
)

func TestCalculateStreak(t *testing.T) {
	now := time.Date(2026, time.August, 15, 14, 0, 0, 0, time.UTC)
	tests := []struct {
		name           string
		executionsDate []time.Time
		now            time.Time
		want           int64
	}{
		{
			name:           "should return streak 0 when there are no executions",
			executionsDate: []time.Time{},
			now:            now,
			want:           0,
		},
		{
			name: "should return streak 1 when executed only today",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 15, 9, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 1,
		},
		{
			name: "should maintain streak at 1 when executed only yesterday",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 14, 18, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 1,
		},
		{
			name: "should return streak 3 when executed two days ago, yesterday, and today",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 13, 12, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 14, 20, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 3,
		},
		{
			name: "should return streak 0 when streak is broken (last execution 2 days ago)",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 12, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 13, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 0,
		},
		{
			name: "should count only current sequence when there is a gap in the past",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 11, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 12, 10, 0, 0, 0, time.UTC),
				// missed day 13
				time.Date(2026, time.August, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 2,
		},
		{
			name: "should count only current sequence when there is a gap in the past 2",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 12, 10, 0, 0, 0, time.UTC),
				// missed day 13
				time.Date(2026, time.August, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 2,
		},
		{
			name: "should not duplicate streak count for multiple executions on the same day",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 14, 9, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 18, 0, 0, 0, time.UTC), // 2nd execution on same day
			},
			now:  now,
			want: 2,
		},
		{
			name: "should calculate extended long streak correctly",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 10, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 11, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 12, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 13, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 14, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 8, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateStreak(tt.executionsDate, tt.now)
			if got != tt.want {
				t.Errorf("CalculateStreak() = %d; expected %d", got, tt.want)
			}
		})
	}
}
