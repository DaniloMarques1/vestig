package cmd

import "testing"

func TestGetDaysLabel(t *testing.T) {
	tests := []struct {
		name   string
		streak int64
		want   string
	}{
		{name: "zero streak", streak: 0, want: "days"},
		{name: "singular streak", streak: 1, want: "day"},
		{name: "plural streak", streak: 2, want: "days"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDaysLabel(tt.streak); got != tt.want {
				t.Errorf("getDaysLabel(%d) = %q, want %q", tt.streak, got, tt.want)
			}
		})
	}
}
